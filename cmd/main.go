package main

import (
	"log"
	"net/http"

	"github.com/juheth/WebDevelopmentWithGo.git/handlers"
)

func main() {
	// Rutas
	http.HandleFunc("/view/", handlers.ViewHandler)
	http.HandleFunc("/edit/", handlers.EditHandler)
	http.HandleFunc("/save/", handlers.SaveHandler)

	// Iniciar el servidor
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
