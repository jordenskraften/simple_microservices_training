package main

import (
	"log"
	"net"
	"test2-micro/transport"
	// Путь к вашему сгенерированному файлу protobuf
)

func main() {
	log.Println("I am the second microservice")

	lis, err := net.Listen("tcp", ":50051") // Слушаем порт 50051 для gRPC
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	transport.StartServer(lis)
}
