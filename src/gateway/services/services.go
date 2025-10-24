package services

import (
	"github.com/!mamvriyskiy/lab1-template/person/model"
	"github.com/!mamvriyskiy/lab1-template/person/repository"
)

type Persons interface {
	GetInfoPerson(personID int) (model.Person, error)
	GetInfoPersons() ([]model.Person, error)
	CreateNewRecordPerson(person model.Person) (model.Person, error)
	UpdateRecordPerson(person model.Person) (model.Person, error)
	DeleteRecordPerson(personID int) error
}

type Services struct {
	Persons
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Persons: NewPersonsService(),
	}
}
