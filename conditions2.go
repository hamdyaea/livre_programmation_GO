package main

import (
    "fmt"
)

func main() {
    age := 40
    //age := 18
    //age := 12

    if age > 18 {
        fmt.Println("Je suis un adulte")
    } else if age == 18 {
        fmt.Println("Je suis juste un adulte !")
    } else {
        fmt.Println("Je ne suis pas un adulte")
    }
}
