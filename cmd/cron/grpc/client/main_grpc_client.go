package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "start/proto/gen"
	"time"
)

func main() {
	credentials := insecure.NewCredentials()

	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(credentials))
	if err != nil {
		log.Fatalf("fialed to open client: %s", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	for {
		req := &pb.GetAllUserRequest{}
		resp, err := client.GetAllUserIDService(context.Background(), req)
		if err != nil {
			log.Printf("Error calling GetAllUserIDService: %v", err)
			// Вы можете выбрать, прервать ли цикл в случае ошибки
			// break
		} else {
			fmt.Printf("User IDs: %v\n", resp.UserIds)
		}

		// Ожидание 10 секунд перед следующим вызовом
		time.Sleep(2 * time.Second)
	}
}
