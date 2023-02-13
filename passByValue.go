package main

import "fmt"

func updater(nombre int) {
    nombre = 5
}

func main() {
    nombre := 10
    updater(nombre)
    fmt.Println(nombre)
}
