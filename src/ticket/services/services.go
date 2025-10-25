package services

import (
	reposiroty "github.com/Mamvriyskiy/lab2-template/src/ticket/reposiroty"
)

type Ticket interface {
}

type Services struct {
	Ticket
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Ticket: NewTicketService(repo.RepoTicket),
	}
}
