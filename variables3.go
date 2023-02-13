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

	fmt.Printf("Votre score est : %v, félicitation.", score)
	fmt.Printf("\nMa couleur est le %v dans cette partie.", couleur)
    fmt.Printf("\nLa finale est : %v, fin.", finale)
}
