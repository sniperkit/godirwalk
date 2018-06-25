package output

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var text Text

func TestTextDoesRegisterItself(t *testing.T) {
	assert.Equal(t, "*output.Text", reflect.TypeOf(ForName("text")).String())
}

func ExampleText_Inline() {
	text.Inline("This is inline")
	// Output: This is inline
}

func ExampleText_Info() {
	text.Info("This is info")
	// Output: This is info
}

func ExampleText_Tick() {
	text.Tick()
	// Output: .
}
