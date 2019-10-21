package oneliner

import (
	"io"
)

type oneliner struct {
	backend io.Writer
}

func (o oneliner) Write(p []byte) (n int, err error) {
	e := getEncoder()
	e.stringBytes(p)
	e.writeNewLine()
	_, err = o.backend.Write(e.getAll())
	if err != nil {
		n = 0
	} else {
		n = len(p)
	}
	putEncoder(e)
	return
}

// Wrap backend, so that every write will be converted to one line string
func Wrap(backend io.Writer) io.Writer {
	return oneliner{
		backend: backend,
	}
}
