package plist

import (
	"io"
	"regexp"

	"github.com/DHowett/go-plist"
)

type Template struct {
}

func (t *Template) Execute(w io.Writer, data interface{}) (err error) {
	b, err := plist.MarshalIndent(data, plist.XMLFormat, "  ")
	if err != nil {
		return
	}

	b = fixBoolTag(b)
	_, err = w.Write(b)

	return
}

var (
	reBoolTag = regexp.MustCompile("<true></true>|<false></false>")
	trueTag   = []byte("<true/>")
	falseTag  = []byte("<false/>")
)

func fixBoolTag(xml []byte) []byte {
	return reBoolTag.ReplaceAllFunc(xml, func(tag []byte) []byte {
		if tag[1] == 't' {
			return trueTag
		}

		return falseTag
	})
}
