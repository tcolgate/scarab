package scarab

import (
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid"

	pb "github.com/tcolgate/scarab/pb"
)

type ProfileRegistration struct {
	ProfileKey
	ProfileWorker
}

type ProfileKey struct {
	Profile string
	Version string
}

type ProfileWorker struct {
	Addr string
	ID   string
	Name string
}
type Profile struct {
	FirstRegistration time.Time
	Args              []*pb.ProfileArg
	Workers           map[ProfileRegistration]struct{}
}

type ProfileRegistry struct {
	sync.RWMutex
	now     func() time.Time
	entries map[ProfileKey]Profile
}

func (pr *ProfileRegistry) Register(req *pb.RegisterProfileRequest) (ProfileRegistration, error) {
	pr.Lock()
	defer pr.Unlock()
	if pr.entries == nil {
		pr.entries = make(map[ProfileKey]Profile)
	}

	key := ProfileKey{
		Profile: req.Profile,
		Version: req.Version,
	}

	t := pr.now()

	prof, ok := pr.entries[key]
	if !ok {
		prof.FirstRegistration = t
	}

	// TODO check the the profile request is compatible/identical
	// with others registered for the same name/version

	r := rand.New(rand.NewSource(t.UnixNano()))
	entropy := ulid.Monotonic(r, 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	pw := ProfileWorker{
		Addr: req.Addr,
		Name: req.Name,
		ID:   id,
	}

	preg := ProfileRegistration{
		ProfileKey:    key,
		ProfileWorker: pw,
	}

	if prof.Workers == nil {
		prof.Workers = make(map[ProfileRegistration]struct{})
	}
	prof.Workers[preg] = struct{}{}
	pr.entries[key] = prof

	return preg, nil
}

func (pr *ProfileRegistry) Unregister(preg ProfileRegistration) {
	pr.Lock()
	defer pr.Unlock()

	if pr.entries == nil {
		return
	}

	key := preg.ProfileKey

	prof := pr.entries[key]
	if prof.Workers == nil {
		return
	}
	delete(prof.Workers, preg)
	pr.entries[key] = prof
}

func (pr *ProfileRegistry) ListProfiles() {
	pr.RLock()
	defer pr.RUnlock()
}
