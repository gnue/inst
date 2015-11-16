package launchdaemon

import (
	"github.com/gnue/inst"
	"github.com/gnue/inst/launchd"
	"github.com/gnue/inst/plist"
)

var Pkg = inst.New(&plist.Template{}, InstallPath)

func Install(name string, data *launchd.Service) (string, error) {
	if data.Label == "" {
		data.Label = name
	}

	return Pkg.Install(name+".plist", 0644, data)
}
