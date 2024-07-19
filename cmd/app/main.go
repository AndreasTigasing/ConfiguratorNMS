package main

import (
	"ConfiguratorNMS/config"
	"ConfiguratorNMS/handlers"
	"fmt"
	"net/http"
)
func main() {
    config.LoadConfig()

	http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/user", handlers.User)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    fmt.Println("Starting server at port 8443")
    if err := http.ListenAndServe(":8443", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}