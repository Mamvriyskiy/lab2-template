package services

import (
	repository "github.com/Mamvriyskiy/lab2-template/src/ticket/repository"
	model "github.com/Mamvriyskiy/lab2-template/src/ticket/model"
)

type Ticket interface {
	GetInfoAboutTiket(ticketUID string) (model.Ticket, error)
	GetInfoAboutTikets(username string) ([]model.Ticket, error)
}

type Services struct {
	Ticket
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Ticket: NewTicketService(repo.RepoTicket),
	}
}
