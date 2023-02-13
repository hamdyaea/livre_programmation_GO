package main

import (
    "net/http"
    "fmt"
)


const port =":5000"


func main() {
    http.HandleFunc("/", Start)
    http.HandleFunc("/Apropos", Apropos)

    fmt.Println("http://127.0.0.1:5000 le serveur à démarré sur le port", port)
    
    http.ListenAndServe(port, nil)
}
