//go:build !darwin && !freebsd && !netbsd && !openbsd && !windows && !js && !wasip1
// +build !darwin,!freebsd,!netbsd,!openbsd,!windows,!js,!wasip1

package term

import (
	"golang.org/x/sys/unix"
)

const (
	getTermios = unix.TCGETS
	setTermios = unix.TCSETS
)
