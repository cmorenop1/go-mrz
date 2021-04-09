# go-mrz
Parse MRZ (Machine Readable Zone) from identity, passport documents.

## Installation

`$ go get github.com/leowilbur/go-mrz

## Example

```go
package main

import (
	"encoding/json"

	"github.com/leowilbur/go-mrz"
)

func main() {
	out, _ := mrz.Parse(`P<USASAAD<<HANY<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	4537163979USA8112060M2407027830274938<837264`)
	buff, _ := json.Marshal(out)
	println(string(buff))
}
```
