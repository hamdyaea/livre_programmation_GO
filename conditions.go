package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())  // Nouvelle clé aléatoire basée sur le temps du système.

    if x := rand.Int(); x%2 == 0 {  // Si le nombre aléatoire placé maintenant la variable x est un nombre paire, test si il y'a un reste au modulo 2 (x%2)
        fmt.Println(x, "nombre paire") // On affiche le nombre paire
    } else {
        fmt.Println(x, "nombre impaire")
    }
}
