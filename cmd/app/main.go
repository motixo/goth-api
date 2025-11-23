package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mot0x0/gopi/internal/adapter/postgres"
	"github.com/mot0x0/gopi/internal/delivery/http"
	"github.com/mot0x0/gopi/internal/usecase/user"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	serverPort := ":" + os.Getenv("SERVER_PORT")
	jwtSecret := os.Getenv("JWT_SECRET")

	db, err := postgres.NewDatabase(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := postgres.NewUserRepository(db.DB)
	usersUC := user.NewUserUsecase(userRepo, jwtSecret)

	server := http.NewServer(usersUC)

	log.Printf("Server starting on port %s", serverPort)
	if err := server.Run(serverPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
