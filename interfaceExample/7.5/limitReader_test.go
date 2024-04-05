package limitReader

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	a := "Hello world"
	reader := strings.NewReader(a)
	lReader := limitReader(reader, 5)

	b := &bytes.Buffer{}
	n, err := b.ReadFrom(lReader)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if n != 5 {
		t.Logf("n=%v", n)
		t.Fail()
	}
	if b.String() != "Hello" {
		t.Logf("%s != %s", b.String(), "Hello")
	}
}