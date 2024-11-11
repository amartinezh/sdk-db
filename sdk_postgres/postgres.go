package sdk_postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostgresSDK struct {
	conn *pgx.Conn
}

func NewPostgresSDK() *PostgresSDK {
	return &PostgresSDK{}
}

func (sdk *PostgresSDK) Connect(connString string) error {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	sdk.conn = conn
	return nil
}

func (sdk *PostgresSDK) Close() error {
	if sdk.conn != nil {
		return sdk.conn.Close(context.Background())
	}
	return nil
}

func (sdk *PostgresSDK) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	if sdk.conn == nil {
		return nil
	}
	return sdk.conn.QueryRow(ctx, query, args...)
}

func (sdk *PostgresSDK) ExecuteQuery(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	if sdk.conn == nil {
		return nil, fmt.Errorf("not connected to the database")
	}
	return sdk.conn.Query(ctx, query, args...)
}

// Exec ejecuta una consulta que no devuelve filas
func (sdk *PostgresSDK) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	if sdk.conn == nil {
		return pgconn.CommandTag{}, fmt.Errorf("not connected to the database")
	}
	return sdk.conn.Exec(ctx, query, args...)
}

// RowsAffected obtiene el número de filas afectadas por la última operación
func RowsAffected(result pgconn.CommandTag) int64 {
	return result.RowsAffected()
}

func (sdk *PostgresSDK) ExecuteTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error {
	if sdk.conn == nil {
		return fmt.Errorf("not connected to the database")
	}

	tx, err := sdk.conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = fn(tx)
	return err
}
