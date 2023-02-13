package handlers

import (
    "text/template"
    "net/http"
    "path/filepath"
    "bytes"
    "programmeweb/config"
    "programmeweb/models"
)

func Start(w http.ResponseWriter, r *http.Request) {
    names := make(map[string]string)
    names["Boss"] = "John"
    renderTemplate(w, "home", &models.TemplateData{
        StringData: names,
        })
}

func Apropos(w http.ResponseWriter, r *http.Request) {
    numb  := make(map[string]int)
    numb["Boss"] = 150

    renderTemplate(w, "about", &models.TemplateData{
        IntData: numb,
        })
}

var appConfig *config.Config

func CreateTemplates(app *config.Config) {
    appConfig = app
}


func renderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
    templateCache := appConfig.TemplateCache

    tmpl, ok := templateCache[tmplName+".page.tmpl"]

    if !ok {
       http.Error(w, "This template is missing", http.StatusInternalServerError)
       return
    }

    buffer := new(bytes.Buffer)
    tmpl.Execute(buffer, td)
    buffer.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
    cache := map[string]*template.Template{}
    pages,err := filepath.Glob("./templates/*.page.tmpl")

    if err != nil {
        return cache, err
    }
    
    for _, page := range pages {
        name := filepath.Base(page)
        tmpl := template.Must(template.ParseFiles(page))
        layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")

        if err != nil {
            return cache, err
        }

        if len(layouts) > 0 {
            tmpl.ParseGlob("./templates/layouts/*.layout.tmpl") 
        }

        cache[name] = tmpl
    }


    return cache, nil
}
