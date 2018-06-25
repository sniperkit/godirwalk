package cli

import (
	"fmt"
	"os"

	// external
	"github.com/pkg/errors"

	// internal
	dirwalk "github.com/sniperkit/snk.golang.dirwalk/pkg"
)

func cleanDirs(dirs ...string) {
	scratchBuffer := make([]byte, 64*1024) // allocate once and re-use each time
	var count, total int
	var err error

	for _, arg := range dirs {
		count, err = pruneEmptyDirectories(arg, scratchBuffer)
		total += count
		if err != nil {
			break
		}
	}

	fmt.Fprintf(os.Stderr, "Removed %d empty directories\n", total)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

func pruneEmptyDirectories(osDirname string, scratchBuffer []byte) (int, error) {
	var count int

	err := dirwalk.Walk(osDirname, &dirwalk.Options{
		Unsorted:      true,
		ScratchBuffer: scratchBuffer,
		Callback: func(_ string, _ *dirwalk.Dirent) error {
			// no-op while diving in; all the fun happens in PostChildrenCallback
			return nil
		},
		PostChildrenCallback: func(osPathname string, _ *dirwalk.Dirent) error {
			deChildren, err := dirwalk.ReadDirents(osPathname, scratchBuffer)
			if err != nil {
				return errors.Wrap(err, "cannot ReadDirents")
			}
			// NOTE: ReadDirents skips "." and ".."
			if len(deChildren) > 0 {
				return nil // this directory has children; no additional work here
			}
			if osPathname == osDirname {
				return nil // do not remove provided root directory
			}
			err = os.Remove(osPathname)
			if err == nil {
				count++
			}
			return err
		},
	})

	return count, err
}
