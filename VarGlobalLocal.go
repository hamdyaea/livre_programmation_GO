package main

import ("fmt")

var x int

func main() {
    x = 5  // Variable Globale
    fmt.Println(x)
    f()
}

func f() {
    x := 10 // Variable Locale
    fmt.Println(x)
}
