package oneliner

import (
	"encoding/json"
	"io"
	"reflect"
	"runtime"
	"unsafe"
)

type oneliner struct {
	backend io.Writer
}

func (o oneliner) Write(p []byte) (n int, err error) {
	// peform unsafe zero-copy conversion from byte slice to string
	// this is safe because of io.Writer contract
	pHeader := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	pString := *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: pHeader.Data,
		Len:  pHeader.Len,
	}))

	err = json.NewEncoder(o.backend).Encode(pString)
	runtime.KeepAlive(p) // make sure p live until here
	if err != nil {
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
