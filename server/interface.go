package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// ServiceRegistrationOption is a type alias for a function that takes a pointer to either a gRPC server or an HTTP gateway server
type ServiceRegistrationOption func(*grpc.Server, *runtime.ServeMux) error

// IAppServer is an interface for a server (gRPC or HTTP gateway)
type IAppServer interface {
	Start(...ServiceRegistrationOption) error
	Stop() error
}
