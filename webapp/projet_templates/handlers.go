package main

import (
    "text/template"
    "net/http"
)

func Start(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "home")
}

func Apropos(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t.Execute(w, nil)
}
