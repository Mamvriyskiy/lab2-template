package services

import (
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/Mamvriyskiy/lab1-template/person/repository"
)

type Ticket interface {
}

type Services struct {
	Ticket
}

func NewServices(repo * repository.Repository) *Services {
	return &Services{
		Ticket: NewTicketService(repo.RepoTicket),
	}
}

