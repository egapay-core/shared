package vault

// ClientEnvironment - the environment of the client
type ClientEnvironment int

const (
	// DevEnvironment - development environment
	DevEnvironment ClientEnvironment = iota
	
	// ProductionEnvironment - production environment
	ProductionEnvironment
)

// DatabaseConnectionConfigurator - the type of database connection string to use
type DatabaseConnectionConfigurator int

const (
	// CustomerConnectionConfigurator - customer database connection string
	CustomerConnectionConfigurator DatabaseConnectionConfigurator = iota
	
	// PaymentConnectionConfigurator - payment database connection string
	PaymentConnectionConfigurator
)
