# copy

file copy install package

* use http.FileSystem interface

## Installation

```sh
$ go get github.com/gnue/inst/pkgs/copy
```

## Usage

```go
import "github.com/gnue/inst/pkgs/copy"
```

## Examples

### Install

```go
package main

import (
	"fmt"
	"github.com/gnue/inst"
	"net/http"
)

func main() {
	fs := copy.New(http.Dir("source"), "/")
	fname, err := fs.Install("target", inst.Local, true)
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
	"net/http"
)

func main() {
	fs := copy.New(http.Dir("source"), "/")
	fname, err := fs.Uninstall("target", inst.Local)
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

