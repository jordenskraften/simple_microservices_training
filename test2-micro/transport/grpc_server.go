package transport

import (
	"log"
	"net"
	pb "test2-micro/generated"

	"google.golang.org/grpc"
)

func StartServer(lis net.Listener) {
	s := grpc.NewServer()

	srv := &MyGrpcServer{}
	pb.RegisterSimpleMessageServiceServer(s, srv)

	addr := lis.Addr().String()
	log.Printf("gRPC Server started on address %s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
