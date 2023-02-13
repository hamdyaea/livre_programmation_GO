package main

import (
    "net/http"
    "fmt"
    "programmeweb/internal/handlers"
)


const port =":5000"


func main() {
    http.HandleFunc("/", handlers.Start)
    http.HandleFunc("/Apropos", handlers.Apropos)

    fmt.Println("http://127.0.0.1:5000 le serveur à démarré sur le port", port)
    
    http.ListenAndServe(port, nil)

}
