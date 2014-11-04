// +build !linux

package statm

// Get returns Statm for given pid along with error. Caller should have access
// to /proc filesystem.
func Get(pid int) (Statm, error) { return Statm{}, ErrNotImplemented }
