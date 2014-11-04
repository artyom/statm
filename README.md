Package statm implements retrieving of memory statistics from
`/proc/[pid]/statm`.

This package is linux-only, calls on other platforms would lead to
`ErrNotImplemented` error.

Install:

	go get -u -v github.com/artyom/statm

Documentation: https://godoc.org/github.com/artyom/statm
