package routes

import (
	"go_studies/web-store-app/controllers"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
