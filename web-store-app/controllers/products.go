package controllers

import (
	"go_studies/web-store-app/models"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SelectProducts()

	templates.ExecuteTemplate(w, "Index", products)
}
