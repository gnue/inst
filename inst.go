package inst

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gnue/merr"
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
	Template        Template
	Locals          []string
	Globals         []string
	InstallAction   func(fname string, loc Locate) error
	UninstallAction func(fname string, loc Locate) error
}

func New(t Template, dirs ...string) *Pkg {
	return &Pkg{Template: t, Locals: dirs}
}

func (pkg *Pkg) InstallPath(loc Locate, mkdir bool) string {
	switch loc {
	case Global:
		return FindDir(pkg.Globals, mkdir)
	default:
		return FindDir(pkg.Locals, mkdir)
	}
}

func (pkg *Pkg) Install(name string, mode os.FileMode, data interface{}, loc Locate, force bool) (fname string, err error) {
	d := pkg.InstallPath(loc, true)
	if d == "" {
		err = fmt.Errorf("inst: no install path")
		return
	}

	fname = filepath.Join(d, name)

	if !force {
		if _, err := os.Lstat(fname); !os.IsNotExist(err) {
			return fname, os.ErrExist
		}
	}

	if data == nil {
		data = name
	}

	if loc == Global {
		var tempDir string
		tempDir, err = ioutil.TempDir("", "inst")
		if err != nil {
			return
		}
		defer os.RemoveAll(tempDir)
		tempFile := filepath.Join(tempDir, name)
		err = pkg.Create(tempFile, mode, data)
		if err != nil {
			return
		}
		err = sudo("cp", tempFile, fname)
	} else {
		err = pkg.Create(fname, mode, data)
	}

	if err != nil {
		return
	}

	if pkg.InstallAction != nil {
		err = pkg.InstallAction(fname, loc)
	}

	return
}

func (pkg *Pkg) Create(fname string, mode os.FileMode, data interface{}) (err error) {
	f, err := os.Create(fname)
	if err != nil {
		return
	}
	defer f.Close()

	err = pkg.Template.Execute(f, data)
	if err != nil {
		return
	}

	err = os.Chmod(fname, mode)

	return
}

func (pkg *Pkg) Uninstall(name string, loc Locate) (fname string, err error) {
	d := pkg.InstallPath(loc, false)
	if d == "" {
		return
	}

	fname = filepath.Join(d, name)
	var errs []error

	if pkg.UninstallAction != nil {
		err := pkg.UninstallAction(fname, loc)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if loc == Global {
		err = sudo("rm", fname)
	} else {
		err = os.Remove(fname)
	}

	if err != nil {
		errs = append(errs, err)
	}
	err = merr.New(errs...)

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

func sudo(arg ...string) error {
	var buf bytes.Buffer

	cmd := exec.Command("sudo", arg...)
	cmd.Stderr = &buf
	err := cmd.Run()
	if err != nil {
		err = errors.New(strings.Trim(buf.String(), "\r\n"))
	}

	return err
}
