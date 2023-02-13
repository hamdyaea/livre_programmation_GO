package main

import ("fmt")

func Bonjour(name string) {
    fmt.Printf("Bonjour, mon prénom est %v.\n", name)
}

func main() {
    Bonjour("Patrick")
}
