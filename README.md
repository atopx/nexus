# Nexus

## Introduction

This is a package for generating unique IDs in a distributed environment. It is based on Twitter's Snowflake algorithm, which has a theoretical maximum of 174 unique IDs, and offers excellent performance. The package is implemented in Go programming language version 1.19.

## Features

- Generates unique IDs in a distributed environment
- Based on Twitter's Snowflake algorithm
- Theoretical maximum of 174 unique IDs
- Excellent performance
- Written in Go programming language version 1.19

## Usage

### Installation
To install the package, run the following command in your terminal:

```
go get github.com/atopx/nexus@latest
```

### Example

```go
package main

import (
	"sync"
	"testing"
	"time"

	"github.com/atopx/nexus"
)

func main() {
	n := nexus.NewNexus(nexus.WithStartTime(time.Now()))
	wg := new(sync.WaitGroup)
	for i := 0; i < 20; i++ {
		go func(n *nexus.Nexus, seq int, wg *sync.WaitGroup) {
            wg.Add(1)
			part, _ := n.NextId()
			defer n.Recover(part)
			t.Logf("nexus part: %+v", part)
			wg.Done()
		}(n, wg)
	}
	wg.Wait()
}
```

## Contributions

We welcome contributions from the community. If you find a bug or have a feature request, please open an issue. If you would like to contribute code, please create a pull request.

## License

This package is released under the MIT License.