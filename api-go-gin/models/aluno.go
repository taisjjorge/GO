package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model //https://gorm.io/docs/models.html#Embedded-Struct
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

//var Alunos []Aluno