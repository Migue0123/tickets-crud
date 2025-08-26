Para ejecutar la aplicación se debe correr el siguiente comando en la termiinal desde la ruta del proyecto
go run main.go

luego para consumir los endpoints se debe ejecutar cada uno añadiendo lal información necesaria. A continuación, algunos ejemplos de cómo consumir los endpoints:

GetAllTickets
curl -X GET http://localhost:8080/tickets

GetTicketByID
curl -X GET http://localhost:8080/tickets/1

CreateTicket
curl -X POST -H "Content-Type: application/json" -d '{"user": "name", "created_date": "26/08/2025", "updated_date": "26/08/2025", "status": "abierto"}' http://localhost:8080/tickets

UpdateTicket
curl -X PUT -H "Content-Type: application/json" -d '{"user": "other name", "created_date": "26/08/2025", "updated_date": "26/08/2025", "status": "cerrado"}' http://localhost:8080/tickets/1

DeleteTicket
curl -X DELETE http://localhost:8080/tickets/1
