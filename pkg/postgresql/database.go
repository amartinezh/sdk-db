package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

// Conectar a la base de datos PostgreSQL
func connectDB(connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return conn, nil
}

// Crear un nuevo registro
func create(conn *pgx.Conn, table string, columns []string, values []interface{}) error {
	colStr := "(" + columns[0]
	valStr := "($1"
	for i := 1; i < len(columns); i++ {
		colStr += ", " + columns[i]
		valStr += fmt.Sprintf(", $%d", i+1)
	}
	colStr += ")"
	valStr += ")"

	query := fmt.Sprintf("INSERT INTO %s %s VALUES %s", table, colStr, valStr)
	_, err := conn.Exec(context.Background(), query, values...)
	return err
}

// Leer registros
func read(conn *pgx.Conn, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := conn.Query(context.Background(), query, args...)
	return rows, err
}

// Actualizar registros
func update(conn *pgx.Conn, table, setClause, whereClause string, args ...interface{}) error {
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, setClause, whereClause)
	_, err := conn.Exec(context.Background(), query, args...)
	return err
}

// Eliminar registros
func delete(conn *pgx.Conn, table, whereClause string, args ...interface{}) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, whereClause)
	_, err := conn.Exec(context.Background(), query, args...)
	return err
}

func main() {
	// Cambia esto a tu cadena de conexión
	connString := "postgres://root:s3rv3r@localhost:5432/crosslist"
	conn, err := connectDB(connString)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer conn.Close(context.Background())

	// Ejemplo de uso:

	// Crear
	err = create(conn, "mi_tabla", []string{"columna1", "columna2"}, []interface{}{"valor1", "valor2"})
	if err != nil {
		log.Printf("Error creating record: %v", err)
	}

	// Leer
	rows, err := read(conn, "SELECT * FROM mi_tabla")
	if err != nil {
		log.Printf("Error reading records: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var columna1, columna2 string
		err := rows.Scan(&columna1, &columna2)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
		}
		fmt.Printf("columna1: %s, columna2: %s\n", columna1, columna2)
	}

	// Actualizar
	err = update(conn, "mi_tabla", "columna1 = $1", "columna2 = $2", "nuevoValor", "valor2")
	if err != nil {
		log.Printf("Error updating record: %v", err)
	}

	// Eliminar
	err = delete(conn, "mi_tabla", "columna1 = $1", "nuevoValor")
	if err != nil {
		log.Printf("Error deleting record: %v", err)
	}
}
