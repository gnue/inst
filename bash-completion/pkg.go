package bash_completion

import (
	"text/template"

	"github.com/gnue/inst"
)

var templ = template.Must(template.New(Name).Parse(Template))
var Pkg = &inst.Pkg{
	Template: templ,
	Locals:   Locals,
	Globals:  Globals,
}

func Install(name string, loc inst.Locate) (string, error) {
	return Pkg.Install(name, 0644, nil, loc)
}
