package inst

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Template interface {
	Execute(wr io.Writer, data interface{}) error
}

type Inst struct {
	Template Template
	Dirs     []string
}

func New(t Template, dirs ...string) *Inst {
	return &Inst{t, dirs}
}

func (inst *Inst) InstallPath() string {
	return FindDir(inst.Dirs, true)
}

func (inst *Inst) Install(name string, mode os.FileMode, data interface{}) (fname string, err error) {
	d := inst.InstallPath()
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

	err = inst.Template.Execute(f, data)
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
