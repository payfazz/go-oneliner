package oneliner

import (
	"encoding/json"
	"io"
)

type oneliner struct {
	backend io.Writer
}

func (o oneliner) Write(p []byte) (n int, err error) {
	data, err := json.Marshal(string(p))
	if err != nil {
		return 0, err
	}
	if len(data) != 0 && data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}
	if _, err := o.backend.Write(data); err != nil {
		return 0, err
	}
	return len(p), nil
}

// Wrap backend, so that every write will be converted to one line string
func Wrap(backend io.Writer) io.Writer {
	return oneliner{
		backend: backend,
	}
}
