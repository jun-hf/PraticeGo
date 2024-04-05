package interfacevalue

import (
	"bytes"
	"io"
)
const debug = true

func Fo(out io.Writer) {
	if out != nil {
		out.Write([]byte("JIJ"))
	}
}

func M(){
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	Fo(buf)
}