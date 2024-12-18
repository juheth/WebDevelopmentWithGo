package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/juheth/WebDevelopmentWithGo.git/models"
	"github.com/juheth/WebDevelopmentWithGo.git/services"
)

var (
	editTemplate *template.Template
	viewTemplate *template.Template
)

func init() {
	var err error
	editTemplate, err = template.ParseFiles("templates/edit.html")
	if err != nil {
		log.Fatalf("Error loading edit template: %v", err)
	}

	viewTemplate, err = template.ParseFiles("templates/view.html")
	if err != nil {
		log.Fatalf("Error loading view template: %v", err)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]

	p, err := services.LoadPage(title)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if err := viewTemplate.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}

	renderTemplates(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, err := services.LoadPage(title)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = editTemplate.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}

	renderTemplates(w, "edit", p)
}

func renderTemplates(w http.ResponseWriter, tmpl string, p *models.Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
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
