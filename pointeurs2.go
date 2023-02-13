package main

import "fmt"

func updater(nombre int) {
    nombre = 3
}

func main() {
    nombre := 10
    fmt.Println(nombre)
    
    fmt.Println("Adresse en mémoire : ",&nombre)

    LePointeur := &nombre
    fmt.Println(LePointeur)

}
