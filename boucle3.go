package main

import ("fmt")


func main() {

    x := 0
    
    for {
        if x > 5 {
            break
        }
        fmt.Println(x)
        x++
    }

}
