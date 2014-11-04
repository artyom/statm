// +build linux

package statm

import (
	"fmt"
	"os"
)

// Get returns Statm for given pid along with error. Caller should have access
// to /proc filesystem.
func Get(pid int) (Statm, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%d/statm", pid))
	if err != nil {
		return Statm{}, err
	}
	defer f.Close()
	s := Statm{}
	n, err := fmt.Fscanf(f, "%d %d %d %d 0 %d 0", &s.Size, &s.Resident, &s.Share, &s.Text, &s.Data)
	if err != nil {
		return Statm{}, err
	}
	if n != 5 {
		return Statm{}, fmt.Errorf("wrong /proc/[pid]/statm format")
	}
	s.Size *= pageSize
	s.Resident *= pageSize
	s.Share *= pageSize
	s.Text *= pageSize
	s.Data *= pageSize
	return s, nil
}

var pageSize uint64

func init() { pageSize = uint64(os.Getpagesize()) }
