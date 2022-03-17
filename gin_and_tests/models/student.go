package models

import (
	"gorm.io/gorm"
	"gopkg.in/validator.v2"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	FederalRegistration  string `json:"federal_registration" validate:"min=9, max=11, regexp=^[0-9]*$"`
	DocumentNumber   string `json:"document_number" validate:"len=9, regexp=^[0-9]{8}[A-Za-z0-9]{1}$"`
}

func(s *Student) Validate() error {
	if err := validator.Validate(s); err != nil {
		return err
	}
	return nil
}
