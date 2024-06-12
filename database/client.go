package database

import (
	"context"
	"fmt"
	"github.com/denisenkom/go-mssqldb/azuread"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	"log"
)

type DbType int

const (
	Cockroach DbType = iota
	Mssql
)

// pool is the connection pool
var (
	pool    *pgxpool.Pool
	mssqlDB *sqlx.DB
)

// Connect creates a new connection pool
func Connect(connString string, dbType DbType) error {
	log.Println("⚙️ connecting to database...")
	var err error
	switch {
	case dbType == Mssql:
		mssqlDB, err = sqlx.Open(azuread.DriverName, connString)
		if err != nil {
			return fmt.Errorf("could not open connection to database: %v", err)
		}
		if err = mssqlDB.Ping(); err != nil {
			return fmt.Errorf("could not ping database: %v", err)
		}
	default:
		pool, err = pgxpool.New(context.Background(), connString)
		if err != nil {
			return fmt.Errorf("could not create connection pool: %v", err)
		}
		if _, err = pool.Exec(context.Background(), "SELECT 1"); err != nil {
			return fmt.Errorf("Failed to ping database: %v\n", err)
		}
	}
	log.Println("🚀 connected to database")
	return nil
}

// GetCockroachDB returns a connection from the pool
func GetCockroachDB() *pgxpool.Conn {
	if pool == nil {
		log.Fatalf("database pool is not initialized properly")
	}
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Fatalf("could not get connection: %v", err)
	}
	return conn
}

// GetMssqlDB returns a connection from the MSSQL database
func GetMssqlDB() *sqlx.DB {
	if mssqlDB == nil {
		log.Fatalf("database pool is not initialized properly")
	}
	return mssqlDB
}
