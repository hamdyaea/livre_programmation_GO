package main

import (
    "net/http"
    "fmt"
)


const port =":5000"

func Start(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bienvenue")
}

func Info(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Informations")
}

func Apropos(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "A propos de ce site web")
}


func main() {
    http.HandleFunc("/", Start)
    http.HandleFunc("/Info", Info)
    http.HandleFunc("/Apropos", Apropos)

    fmt.Println("http://127.0.0.1:5000 le serveur à démarré sur le port", port)
    
    http.ListenAndServe(port, nil)
}
