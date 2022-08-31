package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"strconv"

	"api-rest.com/database"
	"api-rest.com/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Home Page")	
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request){
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
	//json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaPersonalidade(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)
	// for _, personalidade := range models.Personalidades {
	// 	if strconv.Itoa(personalidade.Id) == id {
	// 		json.NewEncoder(w).Encode(personalidade)
	// 	}
	// }
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request){
	var p models.Personalidade
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)

	json.NewEncoder(w).Encode(p)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func AtualizaPersonalidade(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}