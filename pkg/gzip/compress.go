package gzip

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func Compress(inputFileName string, outputFileName string) {
	// open input file
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error while opening file.")
		return
	}
	// close input file
	defer inputFile.Close()

	// create output file
	outputFile, err := os.Create(outputFileName + ".gz")
	if err != nil {
		fmt.Println("Error while creating output file.")
		return
	}
	// close output file
	defer outputFile.Close()

	// create new gzip writer
	gzipWriter := gzip.NewWriter(outputFile)

	// close gzip writer
	defer gzipWriter.Close()

	// copying input files content to output file
	_, err = io.Copy(gzipWriter, inputFile)
	if err != nil {
		fmt.Println("Error while gzip.")
		return
	}

	// final message
	fmt.Println(inputFileName, "gzipped as", outputFileName+".gz")
}
