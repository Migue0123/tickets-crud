package service

import (
	"tickets-crud/internal/model"
	"tickets-crud/internal/store"
)

type TicketService struct {
	store store.TicketStore
}

func New(store store.TicketStore) *TicketService {
	return &TicketService{
		store: store,
	}
}

func (s *TicketService) GetAllTickets() ([]*model.Ticket, error) {
	return s.store.GetAllTickets()
}

func (s *TicketService) GetTicketByID(id int) (*model.Ticket, error) {
	return s.store.GetTicketByID(id)
}

func (s *TicketService) CreateTicket(ticket model.Ticket) (*model.Ticket, error) {
	return s.store.CreateTicket(&ticket)
}

func (s *TicketService) UpdateTicket(id int, ticket model.Ticket) (*model.Ticket, error) {
	return s.store.UpdateTicket(id, &ticket)
}

func (s *TicketService) DeleteTicket(id int) error {
	return s.store.DeleteTicket(id)
}
