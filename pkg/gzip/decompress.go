package gzip

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func Decompress(inputFileName string, outputFileName string) {
	// open input file
	inputFile, err := os.Open(inputFileName + ".gz")
	if err != nil {
		fmt.Println("Error while opening file.")
		return
	}
	// close input file
	defer inputFile.Close()

	// create output file
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error while creating output file.")
		return
	}
	// close output file
	defer outputFile.Close()

	// create new gzip reader
	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		fmt.Println("Error while creating gzip reader.")
		return
	}
	// close gzip reader
	defer gzipReader.Close()

	// copying input files content to output file
	_, err = io.Copy(outputFile, gzipReader)
	if err != nil {
		fmt.Println("Error while decompressing file", inputFileName)
		return
	}

	// final message
	fmt.Println(inputFileName+".gz", "decompressed as", outputFileName)
}
