# CLOG
----- **golang log package** -----

support:
* color output
* log categories
* log levels
* set output destination

## Get Clog
```shell
go get github.com/heejinzzz/clog
```

## Basic Example
```go
package main

import (
	"github.com/heejinzzz/clog"
	"os"
)

func main() {
	// create a new logger
	logger := clog.New(os.Stdout, "{clog} ")

	// print log
	logger.Debug("this is a test")
	logger.Info("this is a test")
	logger.Warn("this is a test")
	logger.Error("this is a test")
	logger.Fatal("this is a test")
}
```