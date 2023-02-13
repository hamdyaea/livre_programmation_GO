// Les generics fonctionnent uniquement à partir de GO version 1.18.
// Vous pouvez voir votre version avec la commande : go version

package main

import "fmt"

func main() {
    fmt.Println(min[float32](10.5,12.7))
    fmt.Println(min[int32] (32,35))
}

func min[T int32 | float32](x, y T) T {
    return x
}
