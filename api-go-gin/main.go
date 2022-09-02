package main

import (
	"github.com/taisjjorge/api-go-gin/database"
	"github.com/taisjjorge/api-go-gin/routes"
)



func main()  {
	database.ConectaComBanco()
	routes.HandleRequests()
}