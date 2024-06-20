package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

// pool is the connection pool
var pool *pgxpool.Pool

// Connect creates a new connection pool
func Connect(connString string) error {
	log.Println("⚙️ connecting to database...")
	ctx := context.Background()
	var err error
	pool, err = pgxpool.New(ctx, connString)
	if err != nil {
		return fmt.Errorf("could not create connection pool: %v", err)
	}
	if _, err = pool.Exec(ctx, "SELECT 1"); err != nil {
		return fmt.Errorf("Failed to ping database: %v\n", err)
	}
	log.Println("🚀 connected to database")
	return nil
}

// GetDB returns a connection from the pool
func GetDB() *pgxpool.Conn {
	if pool == nil {
		log.Fatalf("database pool is not initialized properly")
	}
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("could not get connection: %v", err)
	}
	return conn
}
