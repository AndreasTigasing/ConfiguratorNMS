// handlers/user.go
package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var userTemplate *template.Template

func init() {
    var err error
    userTemplate, err = template.ParseFiles("./templates/base.html", "./templates/user.html")
    if err != nil {
        log.Fatalf("Error parsing templates: %v", err)
    }
}

func User(w http.ResponseWriter, r *http.Request) {
    err := userTemplate.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Println("Error executing template:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
