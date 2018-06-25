// +build dragonfly freebsd openbsd netbsd

package dirwalk

import "syscall"

func direntIno(de *syscall.Dirent) uint64 {
	return uint64(de.Fileno)
}
