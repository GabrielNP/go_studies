package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var templates = template.Must(template.ParseGlob("web-store-app/templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Listening server at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "T-shirt", Description: "Blue, very beaultiful", Price: 39.9, Amount: 1},
		{"Sneakers", "Comfortable", 60, 2},
		{"Headset", "high-end", 200.54, 1},
	}
	templates.ExecuteTemplate(w, "Index", products)
}
