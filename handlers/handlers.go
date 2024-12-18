package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/juheth/WebDevelopmentWithGo.git/models"
	"github.com/juheth/WebDevelopmentWithGo.git/services"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))

func getTitle(w http.ResponseWriter, r *http.Request, prefix string) (string, error) {
	title := r.URL.Path[len(prefix):]
	if !isValidTitle(title) {
		return "", fmt.Errorf("invalid title")
	}
	return title, nil
}

func isValidTitle(title string) bool {
	return !strings.Contains(title, "/") && !strings.Contains(title, "\\")
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r, "/view/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := services.LoadPage(title)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	renderTemplates(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r, "/edit/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := services.LoadPage(title)
	if err != nil {
		p = &models.Page{Title: title, Body: []byte("")}
	}

	renderTemplates(w, "edit", p)
}

func renderTemplates(w http.ResponseWriter, tmpl string, p *models.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("template error %v", err)
	}
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	if title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	body := r.FormValue("body")
	page := &models.Page{Title: title, Body: []byte(body)}

	err := services.SavePage(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error saving page: %v", err)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderError(w http.ResponseWriter, err error, status int) {
	log.Printf("Error: %v", err)
	http.Error(w, err.Error(), status)
}
