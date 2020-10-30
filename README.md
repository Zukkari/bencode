# Bencode

Bencoding library to use for bit torrent protocols.

# Encoding

Stub

# Decoding 

Decoding is done via the exposed `Decode(Reader)` function.

## Integers

```go
package main

import (
    "fmt"
    "strings"
    
    "github.com/Zukkari/bencode"
)

func main() {
    reader := strings.NewReader("i666e")
    decoded, err := bencode.Decode(reader)

    if err != nil {
        fmt.Printf("Error when decoding input: %v\n", err)
    }

    fmt.Println(decoded) // 666
}

```

## Lists

```go
package main

import (
    "fmt"
    "strings"
    
    "github.com/Zukkari/bencode"
)

func main() {
    reader := strings.NewReader("li32ei42ei57ee")
    decoded, err := bencode.Decode(reader)

    if err != nil {
        fmt.Printf("Error when decoding input: %v\n", err)
    }

    fmt.Println(decoded) // [ 32, 42, 57 ]
}

```

## Strings

```go
package main

import (
    "fmt"
    "strings"
    
    "github.com/Zukkari/bencode"
)

func main() {
    reader := strings.NewReader("4:spam")
    decoded, err := bencode.Decode(reader)

    if err != nil {
        fmt.Printf("Error when decoding input: %v\n", err)
    }

    fmt.Println(decoded) // "spam"
}

```

## Dictionaries

```go
package main

import (
    "fmt"
    "strings"
    
    "github.com/Zukkari/bencode"
)

func main() {
    reader := strings.NewReader("d3:bar4:spam3:fooi42ee")
    decoded, err := bencode.Decode(reader)

    if err != nil {
        fmt.Printf("Error when decoding input: %v\n", err)
    }

    fmt.Println(decoded) // "map[bar:spam foo:42]"
}
```