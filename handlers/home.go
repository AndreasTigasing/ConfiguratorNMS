// handlers/home.go
package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var homeTemplate *template.Template

func init() {
    var err error
    homeTemplate, err = template.ParseFiles("templates/base.html", "templates/home.html")
    if err != nil {
        log.Fatalf("Error parsing templates: %v", err)
    }
}

func Home(w http.ResponseWriter, r *http.Request) {
    err := homeTemplate.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Println("Error executing template:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
