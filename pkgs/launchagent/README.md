# launchagent

launchagent install package for macOS

* install path
  * $HOME/Library/LaunchAgents (local)
  * /Library/LaunchAgents (global)

## Installation

```sh
$ go get github.com/gnue/inst/pkgs/launchagent
```

## Usage

```go
import "github.com/gnue/inst/pkgs/launchagent"
```

## Examples

### Install

```go
package main

import (
	"fmt"
	"github.com/gnue/inst"
	"github.com/gnue/inst/launchd"
	"github.com/gnue/inst/pkgs/launchagent"
)

func main() {
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

```

install launchd plist

### Uninstall

```go
package main

import (
	"fmt"
	"github.com/gnue/inst"
	"github.com/gnue/inst/pkgs/launchagent"
)

func main() {
	fname, err := launchagent.Uninstall("label", inst.Local)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("uninstall %q\n", fname)
	}
}

```

uninstall launchd plist

## Author

[gnue](https://github.com/gnue)

## License

[MIT](LICENSE.txt)

