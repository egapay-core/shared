package vault

// GrpcServerConfig - grpc server connection configuration
type GrpcServerConfig struct {
	Host string `env:"GRPC_HOST" env-default:"localhost"`
	Port string `env:"GRPC_PORT" env-default:"50051"`
}

// HttpServerConfig - HTTP server connection configuration
type HttpServerConfig struct {
	Host string `env:"HTTP_HOST" env-default:"localhost"`
	Port string `env:"HTTP_PORT" env-default:"50052"`
}

// PayPartnerServicesConfig - pay partner services configuration
type PayPartnerServicesConfig struct {
	Host string `env:"PAY_PARTNER_SVC_HOST" env-default:"localhost"`
	Port string `env:"PAY_PARTNER_SVC_PORT" env-default:"443"`
}
