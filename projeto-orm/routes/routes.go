package routes

import (
	"log"
	"net/http"

	"api-rest.com/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest(){
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.AtualizaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}