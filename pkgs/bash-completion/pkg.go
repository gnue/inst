// bash-completion install package
//
// * for use "github.com/jessevdk/go-flags" CLI
// * install path
//   * $BASH_COMPLETION_DIR (local)
//   * $BASH_COMPLETION_COMPAT_DIR (global)
package bash_completion

import (
	"text/template"

	"github.com/gnue/inst"
)

var templ = template.Must(template.New("bash").Parse(Template))
var Pkg = &inst.Pkg{
	Template: templ,
	Locals:   Locals,
	Globals:  Globals,
}

func Install(name string, loc inst.Locate, force bool) (string, error) {
	return Pkg.Install(name, 0644, nil, loc, force)
}

func Uninstall(name string, loc inst.Locate) (string, error) {
	return Pkg.Uninstall(name, loc)
}
