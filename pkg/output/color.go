package output

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"

	// internal - core
	config "github.com/sniperkit/snk.golang.dirwalk/pkg/config"
)

const (
	defaultInterval = 300
	minInterval     = 250
	defaultColor    = "yellow"
)

var (
	spin *spinner.Spinner
	cfg  *config.OutputConfig
)

// Color is a color text output
type Color struct {
}

// Configure configures the output
func (c *Color) Configure(oc *config.OutputConfig) {
	cfg = oc
}

// Inline displays text in line
func (c *Color) Inline(s string) {
	fmt.Print(color.GreenString(s))
}

func (c *Color) Interface(m interface{}) {}

// Info displays information
func (c *Color) Info(s string) {
	color.Green(s)
}

// Error displays an error
func (c *Color) Error(s string) {
	color.Red(s)
}

// Fatal displays an error and ends the program
func (c *Color) Fatal(s string) {
	c.Error(s)
	os.Exit(1)
}

// Tick displays evidence that the program is working
func (c *Color) Tick() {
	if spin == nil {
		index := 0
		interval := defaultInterval
		clr := defaultColor
		if cfg != nil {
			index = cfg.SpinnerIndex
			if index < 0 || index > len(spinner.CharSets) {
				index = 0
			}
			interval = cfg.SpinnerInterval
			if interval < minInterval {
				interval = minInterval
			}
			clr = cfg.SpinnerColor
			if clr == "" {
				clr = defaultColor
			}
		}
		spin = spinner.New(spinner.CharSets[index], time.Duration(interval)*time.Millisecond)
		spin.Suffix = color.CyanString(" Updating")
		if err := spin.Color(clr); err != nil {
			c.Error(err.Error())
		}
	}
}

func init() {
	registerOutput(&Color{})
}
