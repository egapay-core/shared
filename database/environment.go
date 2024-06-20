package database

// ConnectionConfigurator - the type of database connection string to use
type ConnectionConfigurator int

const (
	// CustomerConnectionConfigurator - customer database connection string
	CustomerConnectionConfigurator ConnectionConfigurator = iota
	
	// PaymentConnectionConfigurator - payment database connection string
	PaymentConnectionConfigurator
)
