package services

import "github.com/!mamvriyskiy/lab2-template/src/ticket/repository"

type TicketService struct {
	repo repository.RepoTicket
}

func NewTicketService(repo repository.RepoTicket) *TicketService {
	return &TicketService{repo: repo}
}
