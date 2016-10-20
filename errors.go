package inst

import (
	"errors"
	"strings"
)

var (
	ErrNoInstallPath = errors.New("no install path")
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	op := e.Op
	if op == "" {
		op = "inst"
	}

	s := op + " " + e.Path + ": " + e.Err.Error()

	a := EmptyVars(e.Path)
	if 0 < len(a) {
		s += ", please set the environment variable " + strings.Join(a, ", ")
	}

	return s
}
