package transport

import (
	"context"
	"log"

	pb "test2-micro/generated" // Путь к вашему сгенерированному файлу protobuf
	"test2-micro/utils"
)

type MyGrpcServer struct{}

func (s *MyGrpcServer) SimpleMessageMethod(ctx context.Context, request *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	str := request.GetBody()
	log.Printf("Received GRPC message from first microservice: %s", str)

	// Переворачиваем строку
	reversed := utils.ReverseString(str)

	// Создаем объект SimpleMessage с перевернутой строкой и возвращаем его
	finMsg := &pb.SimpleMessage{Body: reversed}
	return finMsg, nil
}
