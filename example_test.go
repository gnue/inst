package inst_test

import (
	"encoding/json"
	"io"

	"github.com/gnue/inst"
)

type Template struct {
}

func (t *Template) Execute(w io.Writer, data interface{}) error {
	e := json.NewEncoder(w)
	return e.Encode(data)
}

// install file
func ExampleInstall() {
	//var templ = template.Must(template.New("config").Parse("name: {{.Name}}\n"))
	var Pkg = &inst.Pkg{
		//Template:        templ,
		Template:        &Template{},
		Locals:          []string{"$XDG_CONFIG_HOME", "$HOME/.config"},
		Globals:         []string{"/usr/local/etc"},
		InstallAction:   nil,
		UninstallAction: nil,
	}

	type Config struct {
		Name string
	}

	Pkg.Install("config.txt", 0644, &Config{Name: "Alice"}, inst.Local, false)
}

// uninstall file
func ExampleUninstall() {
	var Pkg = &inst.Pkg{
		Locals:  []string{"$XDG_CONFIG_HOME", "$HOME/.config"},
		Globals: []string{"/usr/local/etc"},
	}

	Pkg.Uninstall("config.txt", inst.Local)
}
