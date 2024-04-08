package spinner

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100*time.Millisecond)
	fibN := fib(45)
	fmt.Printf("\r %v", fibN)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x -1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}