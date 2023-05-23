package main

import (
	"flag"
	"fmt"
	"runtime"

	compress_gzip "github.com/AlperAkca79/gozip/internal/compress-gzip"
)

const (
	programName = "gozip"

	// edit prgram version manually
	programVersion = "1.0.0"
)

func main() {
	// define command-line flags
	compressFormat := flag.String("f", "gzip", "Compress your file with `gzip` format.")
	inputFile := flag.String("i", "", "Input file.")
	version := flag.Bool("v", false, "Prints version of gozip.")

	// parse command-line flag
	flag.Parse()

	// check the version command-line flag, if true print the gozip version
	if *version {
		fmt.Print(programName, programVersion, " ", GetOS(), "/", GetArch(), "\n")
	}

	outputFile := *inputFile + ".gz"

	// check compress format and compress
	switch *compressFormat {
	case "gzip":
		compress_gzip.Compress(*inputFile, outputFile)
	}
}

// get operating system
func GetOS() string {
	switch runtime.GOOS {
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	case "darwin":
		return "macOS"
	}

	return "undefined-operating-system"
}

// get architecture of operating system
func GetArch() string {
	switch runtime.GOARCH {
	case "386":
		return "386"
	case "amd64":
		return "amd64"
	case "arm":
		return "arm"
	case "arm64":
		return "arm64"
	}

	return "undefined-architecture"
}
