package services

import "github.com/!mamvriyskiy/database_course/main/pkg/repository"

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
