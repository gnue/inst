// file install package
package inst

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gnue/goutils/merr"
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

func (pkg *Pkg) InstallPath(loc Locate, mkdir bool) (string, error) {
	paths := pkg.Locals

	switch loc {
	case Global:
		paths = pkg.Globals
	}

	if len(paths) == 0 {
		return "", &PathError{Err: ErrNoInstallPath}
	}

	d := FindDir(paths, mkdir)
	if d == "" {
		d = paths[0]
		return d, &PathError{Path: d, Err: ErrNoInstallPath}
	}

	return d, nil
}

func (pkg *Pkg) Install(name string, mode os.FileMode, data interface{}, loc Locate, force bool) (fname string, err error) {
	d, err := pkg.InstallPath(loc, true)
	if err != nil {
		return
	}

	fname = filepath.Join(d, name)

	if !force {
		if _, err := os.Lstat(fname); !os.IsNotExist(err) {
			return fname, &os.PathError{"create", name, os.ErrExist}
		}
	}

	if data == nil {
		data = name
	}

	parent := filepath.Dir(fname)

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
		sudo("mkdir", "-p", parent)
		err = sudo("cp", tempFile, fname)
	} else {
		os.MkdirAll(parent, 0755)
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
	d, err := pkg.InstallPath(loc, false)
	if err != nil {
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
		if err == nil || os.IsPermission(err) {
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
