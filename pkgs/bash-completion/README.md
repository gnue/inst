# bash_completion

bash-completion install package

* for use "github.com/jessevdk/go-flags" CLI
* install path
  * $BASH_COMPLETION_DIR (local)
  * $BASH_COMPLETION_COMPAT_DIR (global)

## Installation

```sh
$ go get github.com/gnue/inst/pkgs/bash-completion
```

## Usage

```go
import "github.com/gnue/inst/pkgs/bash-completion"
```

## Examples

### Install

```go
package main

import (
	"fmt"
	"github.com/gnue/inst"
	bash "github.com/gnue/inst/pkgs/bash-completion"
)

func main() {
	fname, err := bash.Install("command", inst.Local, false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("copy %q\n", fname)
	}
}

```

install bash-completion file

### Uninstall

```go
package main

import (
	"fmt"
	"github.com/gnue/inst"
	bash "github.com/gnue/inst/pkgs/bash-completion"
)

func main() {
	fname, err := bash.Uninstall("command", inst.Local)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("remove %q\n", fname)
	}
}

```

uninstall bash-completion file

## Author

[gnue](https://github.com/gnue)

## License

[MIT](LICENSE.txt)

