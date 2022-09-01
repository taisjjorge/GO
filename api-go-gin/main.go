package main

import (
	"github.com/taisjjorge/api-go-gin/database"
	//"github.com/taisjjorge/api-go-gin/models"
	"github.com/taisjjorge/api-go-gin/routes"
)



func main()  {
	database.ConectaComBanco()
	// models.Alunos = []models.Aluno{ 	#### mock alunos
	// 	{Nome: "Tais Amanda", CPF: "12345678910", RG: "202022002"},
	// 	{Nome: "Tanana Tananina", CPF: "98765432110", RG: "555555016"},
	// }
	routes.HandleRequests()
}