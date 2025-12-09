package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
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

type StudentHandler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})
	db.Exec("PRAGMA journal_mode = DELETE;")

	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudente(student Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		return result.Error
	}

	fmt.Println("Create Student!")
	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	students := []Student{}

	err := s.DB.Find(&students).Error
	return students, err

}
