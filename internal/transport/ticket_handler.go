package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tickets-crud/internal/model"
	"tickets-crud/internal/service"
)

type TicketHandler struct {
	service *service.TicketService
}

func New(service *service.TicketService) *TicketHandler {
	return &TicketHandler{
		service: service,
	}
}

func (h *TicketHandler) HandleTickets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tickets, err := h.service.GetAllTickets()
		if err != nil {
			http.Error(w, "Failed to get tickets", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tickets)

	case http.MethodPost:
		var ticket model.Ticket
		if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		created, err := h.service.CreateTicket(ticket)
		if err != nil {
			http.Error(w, "Failed to create ticket", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TicketHandler) HandleTicketByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tickets/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		ticket, err := h.service.GetTicketByID(id)
		if err != nil {
			http.Error(w, "Failed to get ticket", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ticket)

	case http.MethodPut:
		var ticket model.Ticket
		if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		updated, err := h.service.UpdateTicket(id, ticket)
		if err != nil {
			http.Error(w, "Failed to update ticket", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		if err := h.service.DeleteTicket(id); err != nil {
			http.Error(w, "Failed to delete ticket", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
