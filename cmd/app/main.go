// cmd/app/main.go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/amartinezh/sdk-db/pkg/controller"
	"github.com/amartinezh/sdk-db/pkg/service"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

func main() {
	connString := "postgres://postgres:s3rv3r@localhost:5432/crosslist"
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer conn.Close(context.Background())

	personService := &service.PersonService{DB: conn}
	personController := &controller.PersonController{PersonService: personService}

	r := mux.NewRouter()

	// Define las rutas y asócialas con los métodos del controlador
	r.HandleFunc("/persons", personController.AddPerson).Methods("POST")
	r.HandleFunc("/persons", personController.ListPersons).Methods("GET")
	r.HandleFunc("/persons", personController.UpdatePerson).Methods("PUT")
	r.HandleFunc("/persons/{id:[0-9]+}", personController.RemovePerson).Methods("DELETE")

	http.Handle("/", r)
	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
