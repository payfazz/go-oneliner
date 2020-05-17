// Package oneliner for wrapping Write call
package oneliner

import (
	"io"
)

// OneLiner is wrapper type return by Wrap
type OneLiner struct{ io.Writer }

// Write will format p into JSON string and write it
func (o OneLiner) Write(p []byte) (n int, err error) {
	e := getEncoder()
	e.stringBytes(p)
	e.writeNewLine()
	_, err = o.Writer.Write(e.getAll())
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
	// already wrapped?
	if backend, ok := backend.(OneLiner); ok {
		return backend
	}

	return OneLiner{
		Writer: backend,
	}
}
