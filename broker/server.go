package broker

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

// ConsumerServiceRegistrationOption is a type alias for a function that takes a pointer to a proto message
type ConsumerServiceRegistrationOption func(*kafka.Consumer)

// IConsumerServer is an interface for a consumer server
type IConsumerServer interface {
	Start(...ConsumerServiceRegistrationOption) error
	Stop() error
}
