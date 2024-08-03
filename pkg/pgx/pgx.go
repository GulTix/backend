package pgx

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(ctx context.Context, connectionString string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, connectionString)

	if err != nil {
		log.Printf("Unableto to connect to database %+v\n", err)
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
