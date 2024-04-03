package array

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
)

// Write a package that prints the SHA256 hash of it standard input by default but supports a command-line flag to print SHA384 or SHA512 Hash instead


func main() {
	sha := flag.Int("sha", 256, "version of SHA: <usage>: sha=512")
	value := flag.String("val", "", "value to be hashed")
	flag.Parse()
	switch *sha {
	case 256:
		hash := sha256.Sum256([]byte(*value))
		fmt.Printf("%x\n", hash)
	case 512:
		hash := sha512.Sum512([]byte(*value))
		fmt.Printf("%x\n", hash)
	default:
		log.Fatal("Please valid version: 256, 512")
	}



}