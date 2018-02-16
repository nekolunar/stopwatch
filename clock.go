package stopwatch

import (
	"time"
	_ "unsafe" // for go:linkname
)

//go:noescape
//go:linkname nanotime runtime.nanotime
func nanotime() int64

func Start() time.Duration {
	return time.Duration(nanotime())
}

func Stop(t time.Duration) time.Duration {
	return Start() - t
}
