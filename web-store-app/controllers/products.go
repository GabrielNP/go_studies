package controllers

import (
	"go_studies/web-store-app/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SelectProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		p := r.FormValue("price")
		a := r.FormValue("amount")

		price, err := strconv.ParseFloat(p, 64)
		if err != nil {
			log.Println("Error while converting price")
			panic(err.Error())
		}

		amount, err := strconv.Atoi(a)
		if err != nil {
			log.Println("Error while converting amount")
			panic(err.Error())
		}

		models.CreateProduct(name, description, price, amount)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Query().Get("id")
	models.DeleteProduct(product_id)
	http.Redirect(w, r, "/", 301)
}
