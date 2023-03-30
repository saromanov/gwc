package main

import (
	"bufio"
	"flag"
	"log"
	"os"

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
	var stdin []byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = append(stdin, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}

	data := core.New(cfg)
	data.Run(stdin)

}
