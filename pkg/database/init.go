package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func ConnectDB(ctx context.Context, connectionString string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, connectionString)

	if err != nil {
		log.Printf("Unable to to connect to database %+v\n", err)
		os.Exit(1)
	}

	err = conn.Ping(ctx)

	if err != nil {
		log.Printf("Unable to ping database %+v\n", err)
		os.Exit(1)
	}

	log.Printf("Connected to database with connection string %v", connectionString)
	return conn
}

func NewDB(ctx context.Context, connectionString string) DB {
	db := ConnectDB(ctx, connectionString)

	return &dbImpl{
		client: db,
	}
}

func (d *dbImpl) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.client.Query(ctx, query, args...)
}

func (d *dbImpl) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.client.QueryRow(ctx, query, args...)
}

func (d *dbImpl) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.client.Exec(ctx, query, args...)
}
