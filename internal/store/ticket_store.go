package store

import (
	"database/sql"
	"tickets-crud/internal/model"
)

type TicketStore interface {
	GetAllTickets() ([]*model.Ticket, error)
	GetTicketByID(id int) (*model.Ticket, error)
	CreateTicket(ticket *model.Ticket) (*model.Ticket, error)
	UpdateTicket(id int, ticket *model.Ticket) (*model.Ticket, error)
	DeleteTicket(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) TicketStore {
	return &store{db: db}
}

func (s *store) GetAllTickets() ([]*model.Ticket, error) {
	q := `SELECT id, user, created_date, updated_date, status FROM tickets`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*model.Ticket
	for rows.Next() {
		ticket := model.Ticket{}
		if err := rows.Scan(&ticket.ID, &ticket.User, &ticket.CreatedDate, &ticket.UpdatedDate, &ticket.Status); err != nil {
			return nil, err
		}
		tickets = append(tickets, &ticket)
	}
	return tickets, nil
}

func (s *store) GetTicketByID(id int) (*model.Ticket, error) {
	q := `SELECT id, user, created_date, updated_date, status FROM tickets WHERE id = ?`

	ticket := model.Ticket{}
	if err := s.db.QueryRow(q, id).Scan(&ticket.ID, &ticket.User, &ticket.CreatedDate, &ticket.UpdatedDate, &ticket.Status); err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (s *store) CreateTicket(ticket *model.Ticket) (*model.Ticket, error) {
	q := `INSERT INTO tickets (user, created_date, updated_date, status) VALUES (?, ?, ?, ?)`

	resp, err := s.db.Exec(q, ticket.User, ticket.CreatedDate, ticket.UpdatedDate, ticket.Status)
	if err != nil {
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	ticket.ID = int(id)
	return ticket, nil
}

func (s *store) UpdateTicket(id int, ticket *model.Ticket) (*model.Ticket, error) {
	q := `UPDATE tickets SET user = ?, created_date = ?, updated_date = ?, status = ? WHERE id = ?`

	_, err := s.db.Exec(q, ticket.User, ticket.CreatedDate, ticket.UpdatedDate, ticket.Status, ticket.ID)
	if err != nil {
		return nil, err
	}

	ticket.ID = id

	return ticket, nil
}

func (s *store) DeleteTicket(id int) error {
	q := `DELETE FROM tickets WHERE id = ?`

	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}
