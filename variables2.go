package main

import "fmt"

func main() {
	var (
        score int
        nombre int
        finale bool
    )

	score = 5
	nombre = 15
    finale = false

	fmt.Printf("Votre score est : %v, félicitation.", score)
	fmt.Printf("\nLe nombre de joueurs est : %v, dans cette partie.", nombre)
    fmt.Printf("\nLa finale est : %v, fin.", finale)
}
