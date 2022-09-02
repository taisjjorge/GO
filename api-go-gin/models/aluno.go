package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model        //https://gorm.io/docs/models.html#Embedded-Struct
	Nome string `json:"nome" validate:"nonzero"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
}

func ValidadorAlunos(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil{
		return err
	}
	return nil
}
