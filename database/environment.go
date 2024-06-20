package database

// ConnectionConfigurer - the type of database connection string to use
type ConnectionConfigurer int

const (
	// CustomerConnectionConfigurator - customer database connection string
	CustomerConnectionConfigurator ConnectionConfigurer = iota
	
	// PaymentConnectionConfigurator - payment database connection string
	PaymentConnectionConfigurator
)
