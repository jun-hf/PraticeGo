package interfacevalue

import (
	"bytes"
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
	buf.String()

}