// Package oneliner for wrapping Write call
package oneliner

import (
	"io"
)

// OneLiner is wrapper type return by Wrap
type OneLiner struct {
	inner io.Writer
}

// Write will format p into JSON string and write it
func (o OneLiner) Write(p []byte) (n int, err error) {
	e := getEncoder()
	e.stringBytes(p)
	e.writeNewLine()
	_, err = o.inner.Write(e.getAll())
	if err != nil {
		n = 0
	} else {
		n = len(p)
	}
	putEncoder(e)
	return
}

// Wrap inner, so that every write will be converted to one line string
func Wrap(w io.Writer) OneLiner {
	// already wrapped?
	if w, ok := w.(OneLiner); ok {
		return w
	}

	return OneLiner{
		inner: w,
	}
}
