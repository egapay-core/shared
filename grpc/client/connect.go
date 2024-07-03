package client

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// CreateSecureGrpcConnection - creates a secure gRPC connection to the specified address
func CreateSecureGrpcConnection(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	// skip verification of the server's certificate
	tlsCfg := &tls.Config{InsecureSkipVerify: true}

	// create a secure gRPC connection
	creds := grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg))

	// return a connection to the server (or an error)
	return grpc.DialContext(ctx, addr, creds)
}
