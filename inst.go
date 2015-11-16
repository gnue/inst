package inst

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Locate int

const (
	Local = Locate(iota)
	Global
)

type Template interface {
	Execute(wr io.Writer, data interface{}) error
}

type Pkg struct {
	Template Template
	Locals   []string
	Globals  []string
}

func New(t Template, dirs ...string) *Pkg {
	return &Pkg{Template: t, Locals: dirs}
}

func (pkg *Pkg) InstallPath(loc Locate) string {
	switch loc {
	case Global:
		return FindDir(pkg.Globals, true)
	default:
		return FindDir(pkg.Locals, true)
	}
}

func (pkg *Pkg) Install(name string, mode os.FileMode, data interface{}, loc Locate) (fname string, err error) {
	d := pkg.InstallPath(loc)
	if d == "" {
		err = fmt.Errorf("inst: no install path")
		return
	}

	fname = filepath.Join(d, name)
	f, err := os.Create(fname)
	if err != nil {
		return
	}
	defer f.Close()

	if data == nil {
		data = name
	}

	err = pkg.Template.Execute(f, data)
	if err != nil {
		return
	}

	err = os.Chmod(fname, mode)

	return
}

func FindDir(dirs []string, mkdir bool) string {
	mkd := ""

	for _, dir := range dirs {
		d := os.ExpandEnv(dir)

		finfo, err := os.Stat(d)
		if err == nil && finfo.IsDir() {
			return d
		}

		if mkd == "" && os.IsNotExist(err) {
			mkd = d
		}
	}

	if mkdir && mkd != "" {
		err := os.MkdirAll(mkd, 0755)
		if err == nil {
			return mkd
		}
	}

	return ""
}
