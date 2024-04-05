package interfaceExample

import (
	"bufio"
	"strings"
)

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	for scanner.Scan() {
		*l ++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(p), nil
}