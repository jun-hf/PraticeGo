package main

import (
	"bytes"
	"fmt"
	"io"
)
const debug = true

func fo(out io.Writer) {
	if out != nil {
		out.Write([]byte("JIJ"))
	}
}

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	fo(buf)
	fmt.Println(buf.String())
}