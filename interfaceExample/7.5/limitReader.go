package limitReader

import "io"

// write a LimiterReader function that accepts a r and n and report EOF after n
type LimitReader struct {
	reader io.Reader
	count int
	limit int
}

func (l *LimitReader) Read(p []byte) (int, error) {
	c, err := l.reader.Read(p[:l.limit])
	if err != nil {
		return 0, err
	}
	l.count += c
	if l.count >= l.limit {
		return l.count, io.EOF
	}
	return c, nil
}

func limitReader(r io.Reader, n int64) (io.Reader) {
	return &LimitReader{r, 0, int(n)}
}