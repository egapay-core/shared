package cache

import (
	"context"
	"fmt"
	"github.com/egapay-core/shared/vault"
	"github.com/redis/go-redis/v9"
	"log"
)

// defaultDbIndex - default database index
const defaultDbIndex = 0

// cacheConfig - hold the cache configuration details from vault
type cacheConfig struct {
	Host     string `opfield:"server"`
	Port     string `opfield:"port"`
	Password string `opfield:"password"`
}

// RdbConfigurator - redis database configurator
type RdbConfigurator struct {
	ks     *vault.KeyStoreConfig
	client *redis.Client
}

func NewRdbConfigurator(ks *vault.KeyStoreConfig) *RdbConfigurator {
	cfg := &RdbConfigurator{ks: ks}
	if err := loadConfig(cfg, defaultDbIndex); err != nil {
		log.Fatalf("could not load redis configuration: %v", err)
	}
	return cfg
}

// NewConn - creates a new connection to the cache
func (r *RdbConfigurator) NewConn(index int) *redis.Client {
	opts := r.client.Options()
	if opts.DB != index {
		opts.DB = index
	}
	r.client = redis.NewClient(opts)
	return r.client
}

// loadConfig - loads the database configuration from the key vault
func loadConfig(cfg *RdbConfigurator, index int) error {
	// load the cache configuration from the key vault
	var cacheStruct cacheConfig
	if err := cfg.ks.Vault.LoadStructFromItemByTitle(&cacheStruct, vault.CoreCacheConfigItem, cfg.ks.Env.Vault); err != nil {
		log.Printf("failed to load from key vault: %v", err)
		return fmt.Errorf("failed to load from key vault: %v", err)
	}

	// connect to the cache
	addr := fmt.Sprintf("%s:%s", cacheStruct.Host, cacheStruct.Port)
	cfg.client = connect(addr, cacheStruct.Password, index)
	if cfg.client == nil {
		return fmt.Errorf("failed to connect to cache")
	}

	log.Println("🚀 cache configuration loaded successfully")
	return nil
}

// Connect initializes the cache client
func connect(address, password string, db int) *redis.Client {
	log.Println("❄️ connecting to cache...")
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	// check if the connection is successful
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Printf("failed to connect to cache: %v\n", err)
		return nil
	}

	log.Println("🚀️ connected to cache")
	return client
}
