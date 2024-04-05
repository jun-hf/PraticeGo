package interfacevalue

import (
	"fmt"
	"io"
	"os"
)

func E() {
	var w io.Writer = os.Stdout
	w.Write([]byte("Hello"))

	// checking for the type
	fmt.Printf("%T\n", w)
}