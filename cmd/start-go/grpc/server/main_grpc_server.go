package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"start/internal/repository"
	"start/internal/repository/postgres"
	"start/internal/service"
	pb "start/proto/gen"
)

type server struct {
	pb.UserServiceServer
	service *service.Service
}

func (s *server) GetAllUserIDService(ctx context.Context, in *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {

	ids, err := s.service.Todo.GetAllUserIDService()
	if err != nil {
		return nil, status.Error(codes.Unknown, "error GetAllUserIDService in grpc")
	}

	// Преобразуйте полученные идентификаторы в формат, подходящий для ответа gRPC
	userIds := make([]int32, len(ids))
	for i, id := range ids {
		userIds[i] = int32(id)
	}

	// Верните ответ gRPC
	return &pb.GetAllUserResponse{UserIds: userIds}, nil
}

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("error loading environment variable: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Postgres creation error: %s", err.Error())
	}

	ok, err := postgres.CheckDBConnection(db)
	if err != nil {
		log.Fatalf("Connection to the database could not be established: %s", err.Error())
	}
	fmt.Println(ok)

	repos := repository.NewStorageUserPostgres(db)
	service := service.NewServiceUser(repos)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("fialed to listen: %s", err)
	}

	s := grpc.NewServer()
	//инициализирует новый сервер gRPC и регистрирует вашу службу (service) с этим сервером.
	pb.RegisterUserServiceServer(s, &server{service: service})

	fmt.Println("server start on: localhost:50051")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("fialed to serve: %s", err)
	}
}
