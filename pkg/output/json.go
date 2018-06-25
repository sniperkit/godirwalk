package output

import (
	"encoding/json"
	"fmt"
	"os"

	// internal - core
	config "github.com/sniperkit/snk.golang.dirwalk/pkg/config"
)

// Json is a monochrome text output
type Json struct {
}

// Configure no-ops
func (j *Json) Configure(oc *config.OutputConfig) {
}

// Info displays information
func (j *Json) Interface(input interface{}) {
	jsonStr, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		j.Error(err.Error())
	}
	fmt.Println(jsonStr)
}

// Inline displays text in line
func (j *Json) Inline(s string) {
	fmt.Print(j)
}

// Info displays information
func (j *Json) Info(s string) {
	fmt.Println(j)
}

// Error displays an error
func (j *Json) Error(s string) {
	fmt.Fprintln(os.Stderr, j)
}

// Fatal displays an error and ends the program
func (j *Json) Fatal(s string) {
	j.Error(s)
	os.Exit(1)
}

// Tick displays evidence that the program is working
func (j *Json) Tick() {
	fmt.Print(".")
}

func init() {
	registerOutput(&Json{})
}
