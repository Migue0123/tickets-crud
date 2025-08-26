package model

type Ticket struct {
	ID          int    `json:"id"`
	User        string `json:"user"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	Status      string `json:"status"`
}
