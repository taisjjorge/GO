package controllers

import (
	"html/template"
	"net/http"

	"app.com/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosProdutos := models.BuscaProdutos
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}
