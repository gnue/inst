package launchagent_test

import (
	"fmt"

	"github.com/gnue/inst"
	"github.com/gnue/inst/launchd"
	"github.com/gnue/inst/pkgs/launchagent"
)

// install launchd plist
func ExampleInstall() {
	data := &launchd.Service{
		ProgramArguments: []string{"server"},
		KeepAlive:        true,
		RunAtLoad:        true,
	}

	fname, err := launchagent.Install("label", data, inst.Local, false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("install %q\n", fname)
	}
}

// uninstall launchd plist
func ExampleUninstall() {
	fname, err := launchagent.Uninstall("label", inst.Local)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("uninstall %q\n", fname)
	}
}
