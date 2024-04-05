package interfaceExample

import (
	"io"
	"sync"
)

type CountWriter struct {
	m sync.Mutex
	writer io.Writer
	count *int64
}

func (cw *CountWriter) Write(p []byte) (int, error) {
	cw.m.Lock()
	defer cw.m.Unlock()

	c, err := cw.writer.Write(p)
	if err != nil {
		return 0, err
	}
	*cw.count += int64(c)
	return c, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newCount := int64(0)
	var cw *CountWriter = &CountWriter{writer: w, count: &newCount}
	return cw, cw.count
}