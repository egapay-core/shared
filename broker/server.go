package broker

import (
	"google.golang.org/protobuf/proto"
)

// ConsumerServiceRegistrationOption is a type alias for a function that takes a pointer to a proto message
type ConsumerServiceRegistrationOption func(proto.Message)

// IConsumerServer is an interface for a consumer server
type IConsumerServer interface {
	Start(...ConsumerServiceRegistrationOption) error
	Stop() error
}
