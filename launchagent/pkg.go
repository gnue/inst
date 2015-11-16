package launchagent

import (
	"github.com/gnue/inst"
	"github.com/gnue/inst/launchd"
	"github.com/gnue/inst/plist"
)

var Pkg = &inst.Pkg{&plist.Template{}, Locals, Globals}

func Install(name string, data *launchd.Service, loc inst.Locate) (string, error) {
	if data.Label == "" {
		data.Label = name
	}

	return Pkg.Install(name+".plist", 0644, data, loc)
}
