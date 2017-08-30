package launchdaemon_test

import (
	"fmt"

	"github.com/gnue/inst"
	"github.com/gnue/inst/launchd"
	"github.com/gnue/inst/pkgs/launchdaemon"
)

// install launchd plist
func ExampleInstall() {
	data := &launchd.Service{
		ProgramArguments: []string{"server"},
		KeepAlive:        true,
		RunAtLoad:        true,
	}

	fname, err := launchdaemon.Install("label", data, inst.Global, false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("install %q\n", fname)
	}
}

// uninstall launchd plist
func ExampleUninstall() {
	fname, err := launchdaemon.Uninstall("label", inst.Global)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("uninstall %q\n", fname)
	}
}
