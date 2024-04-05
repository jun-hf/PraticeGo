package interfaceExample

import "fmt"

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var a *ByteCounter
	fmt.Fprint(a, "jid")
	fmt.Println(a)
}