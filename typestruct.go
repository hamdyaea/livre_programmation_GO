package main

import "fmt"

type Example struct {
    a string
    b int
    c bool
}

func main() {
    monExample := Example{
        a : "Bonjour",
        b : 10,
        c : true,
    }

fmt.Println(monExample)
}
