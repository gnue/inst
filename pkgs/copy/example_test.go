package copy_test

import (
	"fmt"
	"net/http"

	"github.com/gnue/inst"
	"github.com/gnue/inst/pkgs/copy"
)

// install bash-completion file
func ExampleInstall() {
	fs := copy.New(http.Dir("source"), "/")
	fname, err := fs.Install("target", inst.Local, true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("copy %q\n", fname)
	}
}

// uninstall bash-completion file
func ExampleUninstall() {
	fs := copy.New(http.Dir("source"), "/")
	fname, err := fs.Uninstall("target", inst.Local)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("remove %q\n", fname)
	}
}
