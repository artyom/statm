// +build linux

package statm

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	s, err := Get(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}
	if s.Size == 0 ||
		s.Resident == 0 {
		t.Logf("%+v", s)
		t.Fatal("zero data in statm")
	}
}
