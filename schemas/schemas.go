package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"Name"`
	CPF    int    `json:"CPF"`
	Email  string `json:"Email"`
	Age    int    `json:"Age"`
	Active bool   `json:"Active"`
}

type StudentResponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// DeletedAt time.Time `json:"deletedAt"`
	Name   string `json:"Name"`
	CPF    int    `json:"CPF"`
	Email  string `json:"Email"`
	Age    int    `json:"Age"`
	Active bool   `json:"Active"`
}

func NewResponse(students []Student) []StudentResponse {
	studentsResponse := []StudentResponse{}

	for _, student := range students {
		studentResponse := StudentResponse{
			ID:        int(student.ID),
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			// DeletedAt: student.DeletedAt,
			Name:   student.Name,
			CPF:    student.CPF,
			Email:  student.Email,
			Age:    student.Age,
			Active: student.Active,
		}
		studentsResponse = append(studentsResponse, studentResponse)
	}
	return studentsResponse
}
