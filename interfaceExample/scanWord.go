package interfaceExample

import (
	"bufio"
	"fmt"
	"strings"
)

// Create a new byteCounter type that countes the amount of words

type ByteCounterWord int

func (b *ByteCounterWord) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*b++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(p), nil
}

func e() {
	var w ByteCounterWord
	fmt.Fprint(&w, "Jelo jaijdl ioe")
	fmt.Println(w)
}