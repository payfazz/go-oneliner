package oneliner

import (
	"encoding/json"
	"io"
)

type oneliner struct {
	backend io.Writer
}

func (o oneliner) Write(p []byte) (n int, err error) {
	err = json.NewEncoder(o.backend).Encode(string(p))
	return len(p), err
}

// Wrap backend, so that every write will be converted to one line string
func Wrap(backend io.Writer) io.Writer {
	return oneliner{
		backend: backend,
	}
}
