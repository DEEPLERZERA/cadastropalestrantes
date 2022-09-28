package models

import "gorm.io/gorm"


type Palestrante struct {
	gorm.Model
	Nome string `json:"nome"`
	CPF string `json:"cpf"`
	RG string `json:"rg"`
}