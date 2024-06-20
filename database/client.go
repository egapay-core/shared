package database

import (
	"context"
	"fmt"
	"github.com/egapay-core/shared/vault"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
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

// DbConfigurator - database connection configurator. This struct is used to configure the database connection.
type DbConfigurator struct {
	ks           *vault.KeyStoreConfig
	customerPool *pgxpool.Pool
	paymentPool  *pgxpool.Pool
}

// NewDbConfigurator - creates a new database configurator instance
func NewDbConfigurator(ks *vault.KeyStoreConfig) *DbConfigurator {
	cfg := &DbConfigurator{ks: ks}
	if err := loadConfig(cfg); err != nil {
		log.Fatalf("could not load database configuration: %v", err)
	}
	return cfg
}

// NewConn returns a connection from the pool
func (c *DbConfigurator) NewConn(configurer ConnectionConfigurer) *pgxpool.Conn {
	// get the connection pool
	var pool *pgxpool.Pool
	switch configurer {
	case CustomerConnectionConfigurator:
		pool = c.customerPool
	case PaymentConnectionConfigurator:
		pool = c.paymentPool
	}
	
	if pool == nil {
		log.Printf("could not get connection pool for %v", configurer)
		return nil
	}
	
	// get a connection from the pool
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("could not get connection: %v", err)
	}
	return conn
}

// loadConfig - loads the database configuration from the key vault
func loadConfig(cfg *DbConfigurator) error {
	// function to retrieve the connection pool
	retrievePool := func(ks *vault.KeyStoreConfig,
		configurer ConnectionConfigurer, poolChan chan<- *pgxpool.Pool, wg *sync.WaitGroup) {
		defer wg.Done()
		// get the connection string from the key vault
		connString := retrieveConnectionString(ks, configurer)
		if connString == nil {
			log.Printf("could not get connection string for %v", configurer)
			poolChan <- nil
			return
		}
		
		// connect to the database and return the pool
		poolChan <- connect(*connString)
	}
	
	// wait group to wait for the connection pools to be created
	var wg sync.WaitGroup
	wg.Add(2)
	
	customerPoolChan, paymentPoolChan := make(chan *pgxpool.Pool, 1), make(chan *pgxpool.Pool, 1)
	go retrievePool(cfg.ks, CustomerConnectionConfigurator, customerPoolChan, &wg)
	go retrievePool(cfg.ks, PaymentConnectionConfigurator, paymentPoolChan, &wg)
	wg.Wait()
	
	// assign the connection pools
	cfg.customerPool = <-customerPoolChan
	cfg.paymentPool = <-paymentPoolChan
	
	log.Println("🚀 database configuration loaded successfully")
	return nil
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

// retrieveConnectionString - gets the database connection string.from the key vault for a given ConnectionConfigurer.
func retrieveConnectionString(ks *vault.KeyStoreConfig, configurator ConnectionConfigurer) *string {
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
