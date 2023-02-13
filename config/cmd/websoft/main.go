package main

import (
    "net/http"
    "fmt"
    "programmeweb/internal/handlers"
    "programmeweb/config"
)



func main() {
    
    var appConfig config.Config

    templateCache, err := handlers.CreateTemplateCache()

    if err != nil {
        panic(err)
    }
    
    appConfig.TemplateCache = templateCache
    appConfig.Port = ":5000"
    
    handlers.CreateTemplates(&appConfig)

    http.HandleFunc("/", handlers.Start)
    http.HandleFunc("/Apropos", handlers.Apropos)

    fmt.Println("http://127.0.0.1:5000 le serveur à démarré sur le port", appConfig.Port)
    
    http.ListenAndServe(appConfig.Port, nil)

}
