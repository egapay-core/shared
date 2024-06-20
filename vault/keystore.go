package vault

import (
	"fmt"
	"github.com/1Password/connect-sdk-go/connect"
	"log"
)

// DatabaseConfig - database configuration
type DatabaseConfig struct {
	Host              string `opfield:"server"`
	Port              string `opfield:"port"`
	Username          string `opfield:"username"`
	Password          string `opfield:"password"`
	ConnectionOptions string `opfield:"connection options"`
	DefaultDatabase   string `opfield:"database"`
}

// KeyStoreConfig - keystore configuration
type KeyStoreConfig struct {
	env    *ClientEnvironment
	client connect.Client
}

// NewKeystoreConfig - create a new keystore configuration
func NewKeystoreConfig(env *ClientEnvironment) *KeyStoreConfig {
	ks := &KeyStoreConfig{env: env}
	
	// connect to the vault
	ks.client = connect.NewClient(env.Host, env.Token)
	
	// load credentials from vault
	errChan := make(chan error, 1)
	go loadConfigsFromKeyVault(ks.client, ks, errChan)
	return ks
}

// GetDatabaseConnectionString - gets the database connection string.from the key vault
// for a given DatabaseConnectionConfigurator.
func (ks *KeyStoreConfig) GetDatabaseConnectionString(configurator DatabaseConnectionConfigurator) *string {
	switch configurator {
	case CustomerConnectionConfigurator:
		var dbStruct DatabaseConfig
		if err := ks.client.LoadStructFromItemByTitle(&dbStruct, CustomerDatabaseVaultItem, ks.env.Vault); err != nil {
			log.Printf("failed to load from key vault: %v", err)
			return nil
		}
		return buildConnectionString(&dbStruct)
	case PaymentConnectionConfigurator:
		var dbStruct DatabaseConfig
		if err := ks.client.LoadStructFromItemByTitle(&dbStruct, PaymentDatabaseVaultItem, ks.env.Vault); err != nil {
			log.Printf("failed to load from key vault: %v", err)
			return nil
		}
		return buildConnectionString(&dbStruct)
	default:
		// unknown configurator
		return nil
	}
}

// loadConfigsFromKeyVault loads the sensitive information from the key vault.
func loadConfigsFromKeyVault(client connect.Client, cfg *KeyStoreConfig, errChan chan<- error) {
	defer close(errChan)
	// load the sensitive information from the key vault
	if err := client.LoadStruct(cfg); err != nil {
		log.Printf("failed to load from key vault: %v", err)
		errChan <- err
		return
	}
	
	errChan <- nil
}

// buildConnectionString - builds the connection string for the database.
func buildConnectionString(cfg *DatabaseConfig) *string {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DefaultDatabase, cfg.ConnectionOptions)
	return &connString
}
