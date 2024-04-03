package array

/*
Write a function that counts the number of bits that different from 2 SHA256 hashes
*/
import (
	"crypto/sha256"
	"fmt"
)
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint8) int{
	return int(pc[byte(x)])
}

func compareBitCount(a [32]uint8, b [32]uint8) int {
	var difference int
	for i := range a {
		diff := a[i] ^ b[i]
		difference += PopCount(diff)
	}
	return difference
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(compareBitCount(c1, c2))
}