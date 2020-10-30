package bencode

import (
    "fmt"
    "strings"
)

func main() {
    reader := strings.NewReader("i666e")
    decoded, err := Decode(reader)

    if err != nil {
        fmt.Printf("Error when decoding input: %v\n", err)
    }

    fmt.Println(decoded)
}
