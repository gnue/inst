package plist

import (
	"io"

	"github.com/DHowett/go-plist"
)

type Template struct {
}

func (t *Template) Execute(w io.Writer, data interface{}) (err error) {
	b, err := plist.MarshalIndent(data, plist.XMLFormat, "  ")
	if err != nil {
		return
	}

	_, err = w.Write(b)

	return
}
