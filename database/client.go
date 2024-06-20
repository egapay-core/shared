package database

import (
	"context"
	"fmt"
	"github.com/egapay-core/shared/vault"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

// databaseConfig - database configuration
type databaseConfig struct {
	Host              string `opfield:"server"`
	Port              string `opfield:"port"`
	Username          string `opfield:"username"`
	Password          string `opfield:"password"`
	ConnectionOptions string `opfield:"connection options"`
	DefaultDatabase   string `opfield:"database"`
}

// NewConn returns a connection from the pool
func NewConn(ks *vault.KeyStoreConfig, configurator ConnectionConfigurator) *pgxpool.Conn {
	// get the connection string from the key vault
	connString := retrieveConnectionString(ks, configurator)
	if connString == nil {
		log.Printf("could not get connection string for %v", configurator)
		return nil
	}
	
	// connect to the database
	pool := connect(*connString)
	if pool == nil {
		log.Fatalf("database pool is not initialized properly")
	}
	
	// get a connection from the pool
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("could not get connection: %v", err)
	}
	return conn
}

// connect creates a new connection pool
func connect(connString string) *pgxpool.Pool {
	log.Println("⚙️ connecting to database...")
	ctx := context.Background()
	
	// create a new connection pool
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Printf("could not create connection pool: %v", err)
		return nil
	}
	
	// ping the database to check if the connection is successful
	if _, err = pool.Exec(ctx, "SELECT 1"); err != nil {
		log.Printf("Failed to ping database: %v\n", err)
		return nil
	}
	
	log.Println("🚀 connected to database")
	return pool
}

// retrieveConnectionString - gets the database connection string.from the key vault
// for a given ConnectionConfigurator.
func retrieveConnectionString(ks *vault.KeyStoreConfig, configurator ConnectionConfigurator) *string {
	switch configurator {
	case CustomerConnectionConfigurator:
		var dbStruct databaseConfig
		if err := ks.Vault.LoadStructFromItemByTitle(&dbStruct, vault.CustomerDatabaseVaultItem, ks.Env.Vault); err != nil {
			log.Printf("failed to load from key vault: %v", err)
			return nil
		}
		return buildConnectionString(&dbStruct)
	case PaymentConnectionConfigurator:
		var dbStruct databaseConfig
		if err := ks.Vault.LoadStructFromItemByTitle(&dbStruct, vault.PaymentDatabaseVaultItem, ks.Env.Vault); err != nil {
			log.Printf("failed to load from key vault: %v", err)
			return nil
		}
		return buildConnectionString(&dbStruct)
	default:
		// unknown configurator
		return nil
	}
}

// buildConnectionString - builds the connection string for the database.
func buildConnectionString(cfg *databaseConfig) *string {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DefaultDatabase, cfg.ConnectionOptions)
	return &connString
}
