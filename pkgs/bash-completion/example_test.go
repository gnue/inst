package bash_completion_test

import (
	"fmt"

	"github.com/gnue/inst"
	bash "github.com/gnue/inst/pkgs/bash-completion"
)

// install bash-completion file
func ExampleInstall() {
	fname, err := bash.Install("command", inst.Local, false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("copy %q\n", fname)
	}
}

// uninstall bash-completion file
func ExampleUninstall() {
	fname, err := bash.Uninstall("command", inst.Local)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("remove %q\n", fname)
	}
}
