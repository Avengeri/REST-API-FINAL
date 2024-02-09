package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	srv "start"
	"start/internal/handler"
	"start/internal/repository"
	"start/internal/repository/postgres"
	"start/internal/service"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "start/docs"
)

// @title REST_API_ZAK
// @version 0.0.1
// @description REST API Training Program

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	services := service.NewServiceUser(repos)
	handlers := handler.NewHandler(services)

	server := new(srv.Server)

	port := os.Getenv("SRV_PORT")

	localServer := fmt.Sprintf("the handlergrpc is running on: http://localhost:%s/", port)
	fmt.Println(localServer)

	localPingPong := fmt.Sprintf("ping pong handlergrpc: http://localhost:%s/ping", port)
	fmt.Println(localPingPong)

	localSwag := fmt.Sprintf("swagger: http://localhost:%s/docs/index.html#/", port)
	fmt.Println(localSwag)

	err = server.Run(port, handlers.InitRoutes())
	if err != nil {
		log.Fatalf("the handlergrpc could not be started: %s", err.Error())
	}
}
