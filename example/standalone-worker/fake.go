package main

import (
	"math/rand"
	"time"
)

// sample = NormFloat64() * desiredStdDev + desiredMean

type fake struct {
	rand       *rand.Rand
	hit        float64
	stddev     float64
	hitRatio   float64
	miss       float64
	missStddev float64
}

func newFake(meanHit, hitStddev, missCost, missStddev time.Duration, hitRatio float64) *fake {
	hit := float64(meanHit) / float64(time.Second)
	miss := hit + (float64(missCost) / float64(time.Second))
	return &fake{
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
		hit:        hit,
		miss:       miss,
		stddev:     float64(hitStddev) / float64(time.Second),
		hitRatio:   hitRatio,
		missStddev: float64(missStddev) / float64(time.Second),
	}
}

// Sample generates a duration from a bimodal distribution.
func (f *fake) Sample() (time.Duration, error) {
	mean := f.hit
	stddev := f.stddev

	lucky := f.rand.Float64() < f.hitRatio
	if !lucky {
		mean = f.miss
	}

	samp := f.rand.NormFloat64()*stddev + mean
	if samp < 0 {
		samp = 0.0
	}

	return time.Duration(samp * float64(time.Second)), nil
}
