// main.go

/* JPEG reads a PNG image from the standard input
   and writes it as a JPEG image to the standard output.
   Usage: $ ./mandelbrot | ./go-jpeg > mandelbrot.jpg
   Input format = png */

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"

	_ "image/png" // Register PNG decoder.
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
