# launchdaemon

launchdaemon install package for macOS

* install path
  * $HOME/Library/LaunchAgents (local)
  * /Library/LaunchDaemons (global)

## Installation

```sh
$ go get github.com/gnue/inst/pkgs/launchdaemon
```

## Usage

```go
import "github.com/gnue/inst/pkgs/launchdaemon"
```

## Examples

### Install

```go
package main

import (
	"fmt"
	"github.com/gnue/inst"
	"github.com/gnue/inst/launchd"
	"github.com/gnue/inst/pkgs/launchdaemon"
)

func main() {
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

```

install launchd plist

### Uninstall

```go
package main

import (
	"fmt"
	"github.com/gnue/inst"
	"github.com/gnue/inst/pkgs/launchdaemon"
)

func main() {
	fname, err := launchdaemon.Uninstall("label", inst.Global)
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

