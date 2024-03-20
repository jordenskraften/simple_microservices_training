package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	pb "test1-micro/generated" // Путь к вашему сгенерированному файлу protobuf
	"test1-micro/utils"
	"time"
)

func ReverseHandler(client pb.SimpleMessageServiceClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем URL запроса - если это запрос для иконки сайта, пропускаем его
		if r.URL.Path == "/favicon.ico" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		// Обработка только GET запросов
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		reverseAction(w, r, client)
	}
}

func reverseAction(w http.ResponseWriter, r *http.Request, client pb.SimpleMessageServiceClient) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	ipAddress := r.RemoteAddr
	randomChars := utils.GenerateRandomString(10) //utils.generateRandomString(10)

	message := fmt.Sprintf("%s - %s - %s", currentTime, ipAddress, randomChars)

	response, err := client.SimpleMessageMethod(context.Background(), &pb.SimpleMessage{
		Body: message,
	})
	if err != nil {
		log.Printf("Failed to call SimpleMessageMethod: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respString := response.GetBody()
	log.Printf("Response GRPC message from second microservice: %s", respString)
	fmt.Fprintf(w, respString)
}
