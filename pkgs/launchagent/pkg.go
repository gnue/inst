package launchagent

import (
	"github.com/gnue/inst"
	"github.com/gnue/inst/launchd"
	"github.com/gnue/inst/plist"
)

var Pkg = &inst.Pkg{
	Template:        &plist.Template{},
	Locals:          Locals,
	Globals:         Globals,
	InstallAction:   launchd.InstallAction,
	UninstallAction: launchd.UninstallAction,
}

func Install(name string, data *launchd.Service, loc inst.Locate, force bool) (string, error) {
	if data.Label == "" {
		data.Label = name
	}

	return Pkg.Install(name+".plist", 0644, data, loc, force)
}

func Uninstall(name string, loc inst.Locate) (string, error) {
	return Pkg.Uninstall(name+".plist", loc)
}
