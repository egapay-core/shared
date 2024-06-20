package vault

// ClientEnvironment - the environment of the client
// Host - the host of the client (URL of the client, e.g. https://client.vault.com)
// Token - the token of the client (JWT of the vault)
// Vault - the vault of the client (ID of the vault)
type ClientEnvironment struct {
	Host  string
	Token string
	Vault string
}

// DatabaseConnectionConfigurator - the type of database connection string to use
type DatabaseConnectionConfigurator int

const (
	// CustomerConnectionConfigurator - customer database connection string
	CustomerConnectionConfigurator DatabaseConnectionConfigurator = iota
	
	// PaymentConnectionConfigurator - payment database connection string
	PaymentConnectionConfigurator
)
