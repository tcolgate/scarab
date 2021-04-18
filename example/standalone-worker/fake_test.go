package main

import (
	"testing"
	"time"
)

func TestFake(t *testing.T) {
	f := newFake(
		2*time.Millisecond,
		500*time.Microsecond,
		10*time.Millisecond,
		2*time.Millisecond,
		0.7)

	for i := 0; i < 50; i++ {
		d, _ := f.Sample()
		t.Logf("d: %v", d)
	}
}
