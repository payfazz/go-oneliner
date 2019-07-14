package oneliner

import (
	"bytes"
	"io/ioutil"
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

func Benchmark1(b *testing.B) {
	ol := Wrap(ioutil.Discard)
	data := []byte("test")
	for i := 0; i < b.N; i++ {
		ol.Write(data)
	}
}
