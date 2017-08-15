# SID converter for Go

Converter for Security Identifiers (SID) from binary to string.

## Installation

```
go get github.com/fabiang/sid
```

## Usage

```go
import (
  "fmt"
  "log"

  "github.com/fabiang/sid"
)

func main() {
  mysid := []byte{
    1,
    5,
    0,
    0,
    0,
    0,
    0,
    5,
    21,
    0,
    0,
    0,
    196,
    235,
    38,
    74,
    26,
    193,
    247,
    92,
    104,
    142,
    125,
    166,
    107,
    6,
    0,
    0}

  converted, err := sid.ConvertToString(mysid)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("SID: %s\n", converted)
}
```

## License

[BSD 2-Clause License](LICENSE)
