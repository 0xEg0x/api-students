package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"Name"`
	CPF    int    `json:"CPF"`
	Email  string `json:"Email"`
	Age    int    `json:"Age"`
	Active *bool  `json:"Active"`
}

func errParamRequired(param, typ string) error {
	return fmt.Errorf("parram '%s' of type '%s' is required", param, typ)
}

func (s *StudentRequest) Validate() error {
	if s.Name == "" {
		return errParamRequired("name", "string")
	}
	if s.Email == "" {
		return errParamRequired("Email", "string")
	}
	if s.CPF == 0 {
		return errParamRequired("CPF", "int")
	}
	if s.Age == 0 {
		return errParamRequired("Age", "int")
	}
	if s.Active == nil {
		return errParamRequired("Active", "nil")
	}
	return nil
}
