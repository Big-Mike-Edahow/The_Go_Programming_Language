// methods_test.go

package methods_test

import (
	"strings"
	"time"

	"methods"
)

func ExamplePrintDuration() {
	methods.Print(time.Hour)
}

func ExamplePrintReplacer() {
	methods.Print(new(strings.Replacer))
}
