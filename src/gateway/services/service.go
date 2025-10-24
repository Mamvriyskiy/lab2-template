package services

import (
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/Mamvriyskiy/lab1-template/person/repository"
)

type PersonsService struct {
	repo repository.RepoPersons
}

func NewPersonsService(repo repository.RepoPersons) *PersonsService {
	return &PersonsService{repo: repo}
}

// func (s *PersonsService) GetInfoPerson(personID int) (model.Person, error) {
// 	return s.repo.GetInfoPerson(personID)
// }

// func (s *PersonsService) GetInfoPersons() ([]model.Person, error) {
// 	return s.repo.GetInfoPersons()
// }

// func (s *PersonsService) CreateNewRecordPerson(person model.Person) (model.Person, error) {
// 	return s.repo.CreateNewRecordPerson(person)
// }

// func (s *PersonsService) UpdateRecordPerson(person model.Person) (model.Person, error) {
// 	return s.repo.UpdateRecordPerson(person)
// }

// func (s *PersonsService) DeleteRecordPerson(personID int) error {
// 	return s.repo.DeleteRecordPerson(personID)
// }
