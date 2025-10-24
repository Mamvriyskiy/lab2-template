package services

import (
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/Mamvriyskiy/lab1-template/person/repository"
)

type TicketService struct {
	repo repository.RepoTicket
}

func NewTicketService(repo repository.RepoTicket) *TicketService {
	return &TicketService{repo: repo}
}

