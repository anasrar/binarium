# binarium

Go serialize and deserialize binary struct with reflect.

## Supported Types

- byte
- bool (byte non zero as true)
- uint8, uint16, uint32, uint64
- int8, int16, int32, int64
- float32, float64
- nested struct
- slice
- string

> [!NOTE]
> length `slice` or `string` handle by yourself during serialization / marshal.

## Examples

### Serialization / Marshal

```go
package main

import (
        "encoding/binary"
        "fmt"

        "github.com/anasrar/binarium"
)

type ExampleStruct struct {
        A uint8
        B float32
}

func main() {
        example := ExampleStruct{
                A: 255,
                B: 1.0,
        }

        b, err := binarium.Marshal(binary.BigEndian, example)
        if err != nil {
                fmt.Println(err)
        }

        // [255 63 128 0 0]
        fmt.Println(b)
}
```

### Deserialization / Unmarshal

```go
package main

import (
        "encoding/binary"
        "fmt"

        "github.com/anasrar/binarium"
)

type ExampleStruct struct {
        A uint8
        B float32
}

func main() {
        b := []byte{255, 63, 128, 0, 0}

        example := ExampleStruct{}

        err := binarium.Unmarshal(b, binary.BigEndian, &example)
        if err != nil {
                fmt.Println(err)
        }

        // {255 1}
        fmt.Println(example)
}
```

More examples on `example`.

## Tags

- `length`

  use for `slice` or `string`, pointing to field name as length data.

  ```go
  type ExampleString struct {
        LengthString uint16
        String       string `length:"LengthString"`
  }
  ```

- `skip`

  use for skip field from serialize and deserialize

  ```go
  type ExampleString struct {
        Skip         float64 `skip:""`
        LengthString uint16
        String       string `length:"LengthString"`
  }
  ```

## Background

Use for reverse engineer file format or package sent over socket proxy. Some of them had really wacky structure (I assume for backward compatibility and also not consistent with length type for string and slice), that's why tags for `length` and `skip` exists.
