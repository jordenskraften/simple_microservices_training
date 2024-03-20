package server

import (
	"sync"

	"google.golang.org/grpc"
)

var (
	once     sync.Once
	grpcConn *grpc.ClientConn
	grpcErr  error
)

func GetGRPCConn(addr string) (*grpc.ClientConn, error) {
	once.Do(func() {
		grpcConn, grpcErr = grpc.Dial(addr, grpc.WithInsecure())

	})
	return grpcConn, grpcErr
}
