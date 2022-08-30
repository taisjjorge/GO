package main

import (
	"api-rest.com/models"
	"api-rest.com/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Adriana Varejao", Historia: "Artistona"},
		{Id: 2, Nome: "Helio Oitica", Historia: "Artista"},
	}

	routes.HandleRequest()
}
