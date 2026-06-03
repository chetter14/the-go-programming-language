// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	outputFormat := flag.String("outFmt", "jpeg", "output format of image")
	flag.Parse()

	if err := convert(os.Stdin, os.Stdout, *outputFormat); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, outFmt string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch outFmt {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)
	default:
		return fmt.Errorf("inappropriate output format")
	}
}
