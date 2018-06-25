package output

import (
	"fmt"
	"os"

	// internal - core
	config "github.com/sniperkit/snk.golang.dirwalk/pkg/config"
)

// Text is a monochrome text output
type Text struct {
}

// Configure no-ops
func (t *Text) Configure(oc *config.OutputConfig) {
}

// Inline displays text in line
func (t *Text) Inline(s string) {
	fmt.Print(s)
}

// Print displays an interface{}
func (t *Text) Interface(m interface{}) {}

// Info displays information
func (t *Text) Info(s string) {
	fmt.Println(s)
}

// Error displays an error
func (t *Text) Error(s string) {
	fmt.Fprintln(os.Stderr, s)
}

// Fatal displays an error and ends the program
func (t *Text) Fatal(s string) {
	t.Error(s)
	os.Exit(1)
}

// Tick displays evidence that the program is working
func (t *Text) Tick() {
	fmt.Print(".")
}

func init() {
	registerOutput(&Text{})
}
