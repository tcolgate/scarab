package scarab

import pb "github.com/tcolgate/scarab/pb"

// Server implments the leaders service.
type Server struct{}

func (*Server) Register(*pb.RegisterRequest, pb.Leader_RegisterServer) error {
	return nil
}
