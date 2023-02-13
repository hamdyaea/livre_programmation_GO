package main

import "fmt"

func main() {
    // Définir une variable avec une fonction anonyme
    add := func(a,b int) int {
        return a + b
    }

    // Appeler la fonction anonyme
    result := add(1,2)
    fmt.Println(result)
}
