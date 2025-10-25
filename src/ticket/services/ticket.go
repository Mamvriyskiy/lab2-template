package services

import (
	repository "github.com/Mamvriyskiy/lab2-template/src/ticket/repository"
	model "github.com/Mamvriyskiy/lab2-template/src/ticket/model"
)

type TicketService struct {
	repo repository.RepoTicket
}

func NewTicketService(repo repository.RepoTicket) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) GetInfoAboutTiket(ticketUID string) (model.Ticket, error) {
	return s.repo.GetInfoAboutTiket(ticketUID)
}

func (s *TicketService) GetInfoAboutTikets(username string) ([]model.Ticket, error) {
	return s.repo.GetInfoAboutTikets(username)
}

