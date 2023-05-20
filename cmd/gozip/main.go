package main

import (
	"flag"
	"fmt"

	"github.com/AlperAkca79/gozip/internal/compress-gzip"
)

func main() {
	// define command-line flags
	compressFormat := flag.String("f", "gzip", "Compress your file with `gzip` format.")
	inputFile := flag.String("i", "", "Input file.")

	// parse command-line flag
	flag.Parse()

	// check inputFile command-line flag, show error message if empty
	if *inputFile == "" {
		fmt.Println("Error: Input file cannot be empty. Try again...")
	}

	outputFile := *inputFile + ".gz"

	switch *compressFormat {
	case "gzip":
		compress_gzip.Compress(*inputFile, outputFile)
	}
}