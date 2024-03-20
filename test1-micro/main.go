package main

import (
	"log"
	"net/http"
	"os"
	pb "test1-micro/generated" // Путь к вашему сгенерированному файлу protobuf
	"test1-micro/server"
	"test1-micro/server/handlers"
	"time"
)

func main() {
	log.Println("I am the first microservice")

	grpcAddr := os.Getenv("GRPC_SERVICE_ADDR")
	if grpcAddr == "" {
		log.Fatal("GRPC_SERVICE_ADDR environment variable is not set")
		grpcAddr = "localhost:50051"
	}
	connIsGood := false
	for !connIsGood {
		//до тех пор пока не подключимся
		conn, err := server.GetGRPCConn(grpcAddr)
		if err != nil {
			log.Fatalf("Failed to connect to gRPC server: %v", err)
			time.Sleep(time.Second)
			conn.Close()
			continue
		} else {
			connIsGood = true
			defer conn.Close()
			client := pb.NewSimpleMessageServiceClient(conn)

			http.HandleFunc("/", handlers.ReverseHandler(client))
		}
	}

	httpPort := ":8080"
	log.Printf("Server started on port %s", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
