package main

import (
	"flag"

	"github.com/saromanov/gwc/internal/core"
)

func main() {
	// Define command line flags
	wordsPtr := flag.Bool("w", false, "count of words")
	bytesPtr := flag.Bool("b", false, "count of bytes")

	// Parse command line arguments
	flag.Parse()

	// Access command line flag values
	cfg := core.Config{
		Words: *wordsPtr,
		Bytes: *bytesPtr,
	}

	data := core.New(cfg)
	data.Run()

}
