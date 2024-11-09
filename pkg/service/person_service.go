// pkg/service/person_service.go
package service

import (
	"context"
	"log"

	"github.com/amartinezh/sdk-db/pkg/model"

	"github.com/jackc/pgx/v4"
)

type PersonService struct {
	DB *pgx.Conn
}

// Create a new person
func (ps *PersonService) CreatePerson(person model.Person) error {
	query := "INSERT INTO person (id, name) VALUES ($1, $2)"
	_, err := ps.DB.Exec(context.Background(), query, person.ID, person.Name)
	return err
}

// Get all persons
func (ps *PersonService) GetAllPersons() ([]model.Person, error) {
	query := "SELECT id, name FROM dt.person"
	rows, err := ps.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []model.Person
	for rows.Next() {
		var person model.Person
		if err := rows.Scan(&person.ID, &person.Name); err != nil {
			log.Println("Error scanning person:", err)
			continue
		}
		persons = append(persons, person)
	}
	return persons, nil
}

// Update a person
func (ps *PersonService) UpdatePerson(person model.Person) error {
	query := "UPDATE person SET name=$1 WHERE id=$2"
	_, err := ps.DB.Exec(context.Background(), query, person.Name, person.ID)
	return err
}

// Delete a person
func (ps *PersonService) DeletePerson(personID int) error {
	query := "DELETE FROM person WHERE id=$1"
	_, err := ps.DB.Exec(context.Background(), query, personID)
	return err
}
