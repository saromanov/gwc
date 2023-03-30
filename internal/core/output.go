package core

import "fmt"

func output(r result) {
	fmt.Printf("Bytes: %d\n", r.bytesCount)
	fmt.Printf("Words: %d", r.wordsCount)
}