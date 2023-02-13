package main

import "fmt"

func main() {
	var (
        score int
        couleur string
        finale bool
    )

	score = 5
	couleur = "vert"
    finale = true
    
    fmt.Printf("Votre score est : %v, félicitation. \nMa couleur est le %v dans cette partie. \nLa finale est : %v, fin.", score, couleur, finale)
}
