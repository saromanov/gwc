package main

import (
	"flag"
	"fmt"

	"github.com/saromnov/gwc/internal/core"
)

func main() {
	// Define command line flags
	wordsPtr := flag.Bool("w", false, "count of words")
	bytesPtr := flag.Bool("b", false, "count of bytes")

	// Parse command line arguments
	flag.Parse()

	// Access command line flag values
	cfg := &core.Config{
		wordsPtr: *wordsPtr,
		bytesPtr: *bytesPtr,
	}

	data := core.New(cfg)
	data.Run()

}
