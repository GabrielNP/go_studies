package routes

import (
	"go_studies/web-store-app/controllers"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.Index)
}
