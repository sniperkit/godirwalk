package main

import (
	"fmt"
	"os"

	// internal
	dirwalk "github.com/sniperkit/snk.golang.dirwalk/pkg"
)

func main() {
	dirname := "."
	if len(os.Args) > 1 {
		dirname = os.Args[1]
	}

	err := dirwalk.Walk(dirname, &dirwalk.Options{
		// Unsorted: true, // set true for faster yet non-deterministic enumeration (see godoc)
		Callback: func(osPathname string, de *dirwalk.Dirent) error {
			fmt.Printf("%s %s\n", de.ModeType(), osPathname)
			return nil
		},
		ErrorCallback: func(osPathname string, err error) dirwalk.ErrorAction {
			// Your program may want to log the error somehow.
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)

			// For the purposes of this example, a simple SkipNode will suffice,
			// although in reality perhaps additional logic might be called for.
			return dirwalk.SkipNode
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
