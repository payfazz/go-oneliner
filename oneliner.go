package oneliner

import (
	"io"
)

// OneLiner is wrapper type return by Wrap
type OneLiner struct {
	backend io.Writer
}

func (o OneLiner) Write(p []byte) (n int, err error) {
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
func Wrap(backend io.Writer) OneLiner {
	return OneLiner{
		backend: backend,
	}
}
