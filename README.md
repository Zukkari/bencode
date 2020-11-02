# Bencode

Bencoding library to use for bit torrent protocols.

# Encoding

Encoding is done via the exposed `Encode(interface{})` function

```go
package main

import (
    "fmt"
    "strings"
    
    "github.com/Zukkari/bencode"
)

func main() {
    // Int

    encodedInt, err := bencode.Encode(32)
    if err != nil {
        fmt.Printf("Got error when encoding: %v\n", err)
    }

    fmt.Println(encodedInt) // i32e

    // String

    encodedString, err := bencode.Encode("test")
    if err != nil {
        fmt.Printf("Got error when encoding: %v\n", err)
    }

    fmt.Println(encodedString) // 4:test

    // List

    inputSlice := make([]interface{}, 3)
    inputSlice[0] = 1
    inputSlice[1] = 2
    inputSlice[2] = 3

    encodedList, err := bencode.Encode(inputSlice)
    if err != nil {
        fmt.Printf("Got error when encoding: %v\n", err)
    }

    fmt.Println(encodedList) // li1ei2ei3ee

    // Dict
    inputDict := make(map[string]interface{})
    inputDict["a"] = 42
    inputDict["b"] = "spam"

    arr := make([]interface{}, 2)
    arr[0] = 42
    arr[1] = "eggs"

    inputDict["c"] = arr

    encodedDict, err := bencode.Encode(inputDict)
    if err != nil {
        fmt.Printf("Got error when encoding: %v\n", err)
    }

    fmt.Println(encodedDict) // d1:ai42e1:b4:spam1:cli42e4:eggsee
}
```

# Decoding 

Decoding is done via the exposed `Decode(Reader)` function.

```go
package main

import (
    "fmt"
    "strings"
    
    "github.com/Zukkari/bencode"
)

func main() {
    // Int

    intReader := strings.NewReader("i666e")
    decodedInt, err := bencode.Decode(intReader)

    if err != nil {
        fmt.Printf("Error when decoding input: %v\n", err)
    }

    fmt.Println(decodedInt) // 666

    // String

    stringReader := strings.NewReader("4:spam")
    decodedString, err := bencode.Decode(stringReader)

    if err != nil {
        fmt.Printf("Error when decoding input: %v\n", err)
    }

    fmt.Println(decodedString) // "spam"

    // List

    listReader := strings.NewReader("li32ei42ei57ee")
    decodedList, err := bencode.Decode(listReader)

    if err != nil {
        fmt.Printf("Error when decoding input: %v\n", err)
    }

    fmt.Println(decodedList) // [ 32, 42, 57 ]

    // Dictionary

    dictReader := strings.NewReader("d3:bar4:spam3:fooi42ee")
    decodedDict, err := bencode.Decode(dictReader)

    if err != nil {
        fmt.Printf("Error when decoding input: %v\n", err)
    }

    fmt.Println(decodedDict) // "map[bar:spam foo:42]"
}
```