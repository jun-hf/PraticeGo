package interfaceExample

import "io"

type couterWritter struct {
	count int64
	writer io.Writer
}

func (cw *couterWritter)Write(p []byte) (int, error) {
	currentCount, err := cw.writer.Write(p)
	if err != nil {
		return currentCount, err
	}
	cw.count += int64(currentCount)
	return currentCount, nil
}

func countingWriter(w io.Writer) (io.Writer, *int64) {
	cw := couterWritter{writer: w, count: 0}
	return &cw, &cw.count
}