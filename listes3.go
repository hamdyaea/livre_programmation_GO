package main

import("fmt")

func main() {
    
    GrandeListe := [...]int{40, 50, 65, 470, 6464677, 5354353532}

    for k,v := range GrandeListe {
        fmt.Printf("Position %d est égal à %d.\n", k, v)
    }
}
