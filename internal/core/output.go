package core

import "fmt"

func output(r result) {
	fmt.Printf("Bytes: %d\n", r.bytesCount)
	fmt.Printf("Words: %d\n", r.wordsCount)
	fmt.Printf("Chars: %d\n", r.charsCount)
	fmt.Printf("Lines: %d\n", r.linesCount)
}