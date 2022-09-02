package main

import (
	//"fmt"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/taisjjorge/api-go-gin/controllers"
	"github.com/taisjjorge/api-go-gin/database"
	"github.com/taisjjorge/api-go-gin/models"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //simplifica resultados test
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock()  {
	aluno := models.Aluno{Nome: "Nome Aluno Teste", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock()  {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeGreetings(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/tais", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	// if res.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d",
	// 		res.Code, http.StatusOK)
	// }

	// ## com assert (go get github.com/stretchr/testify)
	assert.Equal(t, http.StatusOK, res.Code, "Flopou!! Deveriam ser iguais")
	mockDaResposta := `{"greetings!":"Hey tais, joia?"}`
	respostaBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
	//fmt.Println(string(respostaBody))
	//fmt.Println(mockDaResposta)
}

func TestListandoAlunosHandler(t *testing.T) {
	database.ConectaComBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	//fmt.Println(res.Body)
}

func TestBuscaPorCPFHanlder(t *testing.T) {
	database.ConectaComBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678910", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	database.ConectaComBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoId)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoMock models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)
	assert.Equal(t, "Nome Aluno Teste", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12345678910", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBanco()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestEditaUmAlunoHandler(t *testing.T) {
	database.ConectaComBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome Aluno Teste", CPF: "47123456789", RG: "123456700"}
	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome Aluno Teste", alunoMockAtualizado.Nome)
}