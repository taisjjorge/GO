package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taisjjorge/api-go-gin/controllers"
)

func HandleRequests(){
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)
	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run()
}