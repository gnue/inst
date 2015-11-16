package bash_completion

import (
	"text/template"

	"github.com/gnue/inst"
)

var templ = template.Must(template.New(Name).Parse(Template))
var Pkg = inst.New(templ, InstallPath)

func Install(name string) (string, error) {
	return Pkg.Install(name, 0644, nil)
}
