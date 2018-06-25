// +build darwin linux

package dirwalk

import "syscall"

func direntIno(de *syscall.Dirent) uint64 {
	return de.Ino
}
