package main

import (
	"log"
	"net/http"
	"time"

	"github.com/juheth/WebDevelopmentWithGo.git/handlers"
)

type Config struct {
	Port     string
	DataDir  string
	LogLevel string
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

func main() {
	config := Config{
		Port:     ":8080",
		DataDir:  "data",
		LogLevel: "info",
	}

	// Rutas
	http.HandleFunc("/view/", loggingMiddleware(handlers.ViewHandler))
	http.HandleFunc("/edit/", loggingMiddleware(handlers.EditHandler))
	http.HandleFunc("/save/", loggingMiddleware(handlers.SaveHandler))

	// Iniciar el servidor
	log.Printf("Server running on http://localhost%s", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, nil))
}
