package oneliner

import (
	"bytes"
	"testing"
)

func Test1(t *testing.T) {
	b := &bytes.Buffer{}
	ol := Wrap(b)
	ol.Write([]byte("a\nb"))
	if !bytes.Equal(b.Bytes(), []byte("\"a\\nb\"\n")) {
		t.FailNow()
	}
}
