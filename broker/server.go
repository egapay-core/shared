package broker

import "github.com/egapay-core/shared/vault"

// ConsumerServiceRegistrationOption is a type alias for a function that takes a key store configuration,
// a groupID and a list of topics and returns an error
type ConsumerServiceRegistrationOption func(vault.KeyStoreConfig, string, []string) error

// IConsumerServer is an interface for a consumer server
type IConsumerServer interface {
	Start(...ConsumerServiceRegistrationOption) error
	Stop() error
}
