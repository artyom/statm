// Package statm implements retrieving of memory statistics from
// /proc/[pid]/statm
//
// This package is linux-only, calls on other platforms would lead to
// ErrNotImplemented.
package statm // import "github.com/artyom/statm"

import "errors"

// Statm provides information about process memory usage, measured in bytes
type Statm struct {
	Size     uint64 // total program size in bytes (same as VmSize in /proc/[pid]/status)
	Resident uint64 // resident set size in bytes (same as VmRSS in /proc/[pid]/status)
	Share    uint64 // shared size in bytes (from shared mappings)
	Text     uint64 // text (code) in bytes
	Data     uint64 // data + stack in bytes
}

var ErrNotImplemented = errors.New("not implemented")
