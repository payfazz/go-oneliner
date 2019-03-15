package oneliner

import (
	"bytes"
	"encoding/json"
	"io"
)

type oneliner struct {
	backend io.Writer
	buf     *bytes.Buffer
	encoder *json.Encoder
}

func (o *oneliner) Write(p []byte) (n int, err error) {
	o.buf.Reset()
	o.encoder.Encode(string(p))
	_, err = o.backend.Write(o.buf.Bytes())
	return len(p), err
}

// Wrap backend, so that every write will be converted to one line string
func Wrap(backend io.Writer) io.Writer {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	return &oneliner{
		backend: backend,
		buf:     buf,
		encoder: encoder,
	}
}
