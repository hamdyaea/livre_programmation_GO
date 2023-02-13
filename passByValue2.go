package main

import "fmt"

func updater(nombre int) int{
    nombre = 5
    return nombre
}

func main() {
    nombre := 10
    nombre = updater(nombre)
    fmt.Println(nombre)
}
