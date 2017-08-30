# inst

file install package

## Installation

```sh
$ go get github.com/gnue/inst
```

## Usage

```go
import "github.com/gnue/inst"
```

## Examples

### Install

```go
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
```

install file

### Uninstall

```go
package main

import (
	"github.com/gnue/inst"
)

func main() {
	var Pkg = &inst.Pkg{
		Locals:  []string{"$XDG_CONFIG_HOME", "$HOME/.config"},
		Globals: []string{"/usr/local/etc"},
	}

	Pkg.Uninstall("config.txt", inst.Local)
}

```

uninstall file

## Author

[gnue](https://github.com/gnue)

## License

[MIT](LICENSE.txt)

