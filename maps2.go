package main

import "fmt"


func main() {

    PrixDesArticles := map[string]int{
        "boisson" : 10,
        "viande"  :  7,
        "poisson" : 14,
    }


    fmt.Println(PrixDesArticles["poisson"])

    for key, value := range PrixDesArticles {
        fmt.Println(key, value)
    }
}
