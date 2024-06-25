package vault

import (
	"github.com/1Password/connect-sdk-go/connect"
	"log"
)

// KeyStoreConfig - keystore configuration
type KeyStoreConfig struct {
	Env        *ClientEnvironment
	Vault      connect.Client
	GrpcServer *GrpcServerConfig
	HttpServer *HttpServerConfig
}

// NewKeystoreConfig - create a new keystore configuration
func NewKeystoreConfig(env *ClientEnvironment, grpcServer *GrpcServerConfig, httpServer *HttpServerConfig) *KeyStoreConfig {
	ks := &KeyStoreConfig{Env: env, GrpcServer: grpcServer, HttpServer: httpServer}

	// connect to the vault
	ks.Vault = connect.NewClient(env.Host, env.Token)

	// load credentials from vault
	errChan := make(chan error, 1)
	go loadConfigsFromKeyVault(ks.Vault, ks, errChan)
	return ks
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
