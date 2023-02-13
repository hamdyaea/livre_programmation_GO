package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1
    var guess int
    for {
        fmt.Print("Devinez le nombre: ")
        fmt.Scanf("%d\n", &guess)
        if guess < target {
            fmt.Println("Le nombre est plus élevé.")
        } else if guess > target {
            fmt.Println("Le nombre est plus bas.")
        } else {
            fmt.Println("Bravo, vous avez deviné le nombre!")
            break
        }
    }
}

