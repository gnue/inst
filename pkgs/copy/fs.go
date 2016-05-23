package copy

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gnue/httpfs/fsutil"
	"github.com/gnue/inst"
	"github.com/gnue/merr"
)

type FileSystem struct {
	http.FileSystem
	root string
}

func New(fs http.FileSystem, root string) *FileSystem {
	return &FileSystem{fs, root}
}

func (fs *FileSystem) Open(name string) (http.File, error) {
	fname := filepath.Join(fs.root, name)
	return fs.FileSystem.Open(fname)
}

func (fs *FileSystem) Execute(w io.Writer, data interface{}) (err error) {
	fname, ok := data.(string)
	if !ok {
		return os.ErrExist
	}

	f, err := fs.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(w, f)

	return
}

func (fs *FileSystem) Install(dir string, loc inst.Locate, force bool) (string, error) {
	var pkg = &inst.Pkg{
		Template: fs,
		Locals:   []string{dir},
		Globals:  []string{dir},
	}

	var errs []error

	f := fsutil.FileSystem{fs}
	err := f.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if _, err := pkg.Install(path, info.Mode()&os.ModePerm, nil, loc, force); err != nil {
			errs = append(errs, err)
		}

		return nil
	})
	if err != nil {
		errs = append(errs, err)
	}

	return dir, merr.New(errs...)
}

func (fs *FileSystem) Uninstall(dir string, loc inst.Locate) (string, error) {
	var pkg = &inst.Pkg{
		Locals:  []string{dir},
		Globals: []string{dir},
	}

	var errs []error

	f := fsutil.FileSystem{fs}
	err := f.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if _, err := pkg.Uninstall(path, loc); err != nil {
			errs = append(errs, err)
		}

		return nil
	})
	if err != nil {
		errs = append(errs, err)
	}

	return dir, merr.New(errs...)
}
