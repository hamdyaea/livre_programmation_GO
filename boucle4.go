package main

import ("fmt")


func main() {

    x := 0
    
    for ; x <= 12; x++ {
        if x%2 == 1 {
            continue
        }
        fmt.Println(x)
    }
}
