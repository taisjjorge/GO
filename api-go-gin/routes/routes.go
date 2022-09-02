package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taisjjorge/api-go-gin/controllers"
)

func HandleRequests(){
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeAlunos)
	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/home", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.Rota404)
	r.Run()
}