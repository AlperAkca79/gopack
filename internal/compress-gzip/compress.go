package compress_gzip

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

func Compress(inputFile string, outputFile string) {
	// open input file
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error while opening input file.")
	}

	// create `outputFile` writer
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error while creating output file.")
	}

	// create `gzip writer`
	gzipWriter := gzip.NewWriter(output)
	if err != nil {
		fmt.Println("Error while creating output file.")
	}

	_, err = io.Copy(gzipWriter, input)
	if err != nil {
		fmt.Println("Error while copying input file to output file.")
	}

	// close `gzipWriter`
	gzipWriter.Close()

	// close `input`
	input.Close()

	// close `output`
	output.Close()

	// set foreground
	color.Set(color.FgGreen)

	fmt.Print("`gzip` ")

	// unset color
	color.Unset()

	fmt.Println("compression completed succesfully.")
}