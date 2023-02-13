// Les generics fonctionnent uniquement à partir de GO version 1.18.
// Vous pouvez voir votre version avec la commande : go version

package main

import "fmt"

func main() {
    type f float64
    var ref f = 5.12
    fmt.Println(min(ref, 0.4))
}

func min[T ~float64](x, y T) T {
    return x
}
