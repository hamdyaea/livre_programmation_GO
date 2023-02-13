package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// Déclarez une structure pour stocker les données JSON
type Person struct {
    Name string
    Age  int
}

func main() {
    // Créez une instance de la structure Person
    p := Person{Name: "John Doe", Age: 30}

    // Encodez la structure en JSON
    b, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Affichez les données JSON encodées
    fmt.Println(string(b))

    // Écrivez les données JSON encodées dans un fichier
    f, err := os.Create("person.json")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer f.Close()
    _, err = f.Write(b)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Lisez les données JSON à partir d'un fichier
    f, err = os.Open("person.json")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer f.Close()
    var p2 Person
    err = json.NewDecoder(f).Decode(&p2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Affichez les données JSON décodées
    fmt.Println(p2)
}

