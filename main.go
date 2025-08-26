package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"tickets-crud/internal/service"
	"tickets-crud/internal/store"
	"tickets-crud/internal/transport"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./tickets.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := `CREATE TABLE IF NOT EXISTS tickets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user TEXT NOT NULL,
		created_date TEXT NOT NULL,
		updated_date TEXT NOT NULL,
		status TEXT NOT NULL
	);`

	if _, err = db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	//Inyección de dependencias
	ticketStore := store.New(db)
	ticketService := service.New(ticketStore)
	ticketHandler := transport.New(ticketService)

	//Configuración de rutas
	http.HandleFunc("/tickets", ticketHandler.HandleTickets)
	http.HandleFunc("/tickets/", ticketHandler.HandleTicketByID)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
