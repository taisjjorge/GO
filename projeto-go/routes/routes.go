package routes

import (
	"net/http"
	
	"app.com/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}