package db

import (
	"context"
	"database/sql"
	"os"
	// TODO: import a PostgreSQL driver, e.g.
	//   _ "github.com/jackc/pgx/v5/stdlib" (driver name "pgx")
	//   _ "github.com/lib/pq"              (driver name "postgres")
	// or swap this package out for an ORM / query builder of your choice.
)

const defaultDSN = "postgres://jobsuser:jobspass@localhost:5432/jobsdb?sslmode=disable"

// Connect opens a connection pool to PostgreSQL and verifies it with a ping.
// The DSN is read from DATABASE_URL, falling back to the compose.yml defaults.
func Connect(ctx context.Context) (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = defaultDSN
	}

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err := conn.PingContext(ctx); err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}
