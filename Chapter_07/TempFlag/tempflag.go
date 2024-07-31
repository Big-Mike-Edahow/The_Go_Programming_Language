// main.go

/* Tempflag prints the value of its -temp (temperature) flag.
   Usage: $ go run . -temp 32°F */

package main

import (
	"flag"
	"fmt"

	"tempflag/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
