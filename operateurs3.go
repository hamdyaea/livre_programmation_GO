package main

import "fmt"

func main(){
    var (
        x int
        y int
    )

    x = 30
    y = 5

    fmt.Println(x == y)
    fmt.Println(x != y)
    fmt.Println(x < y)
    fmt.Println(x <= y)
    fmt.Println(x > y)
    fmt.Println(x >= y)

}
