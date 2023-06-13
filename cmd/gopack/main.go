package main

import (
	"flag"
	"github.com/AlperAkca79/gopack/pkg/gzip"
)

func main() {
	inputFile := flag.String("i", "", "Set input file.")
	outputFile := flag.String("o", "", "Set output file.")

	flag.Parse()

	gzip.Decompress(*inputFile, *outputFile)
}
