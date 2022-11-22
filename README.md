[![Go Reference](https://img.shields.io/badge/go-reference-%2300ADD8?style=flat-square)](https://pkg.go.dev/github.com/backdround/go-wrappingerror)
[![Tests](https://img.shields.io/github/workflow/status/backdround/go-wrappingerror/tests?label=tests&style=flat-square)](https://github.com/backdround/go-wrappingerror/actions)
[![Codecov](https://img.shields.io/codecov/c/github/backdround/go-wrappingerror?style=flat-square)](https://app.codecov.io/gh/backdround/go-wrappingerror/)
[![Go Report](https://goreportcard.com/badge/github.com/backdround/go-wrappingerror?style=flat-square)](https://goreportcard.com/report/github.com/backdround/go-wrappingerror)

# Wrapping error
Package provides error that can wrap other error and returns pretty error message.


### Istallation
```bash
go get github.com/backdround/go-wrappingerror
```

### Example

```go
package main

import (
	"github.com/backdround/go-wrappingerror"
	"errors"
	"fmt"
	"os"
)

var unableToCreateFile = wrappingerror.NewWrappingError("unable to create file")

// createFile returns internal error
func createFile() error {
	  return errors.New("access denied")
}


// doSomething returns wrapping error.
func doSomething() error {
	err := createFile()
	if err != nil {
		return unableToCreateFile.Wrap(err)
	}

	return nil
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

Output:
```
unable to create file:
  access denied
```
