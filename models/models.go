package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
