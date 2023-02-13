// Les generics fonctionnent uniquement à partir de GO version 1.18.
// Vous pouvez voir votre version avec la commande : go version

package main

import "fmt"

type Nombre interface {
    int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | ~float64
}

func main() {
    type f float64
    var ref f = 5.12
    fmt.Println(min(ref, 0.4))
}

func min[T Nombre](x, y T) T {
    if x < y {
        return x
    }
    return y
}
