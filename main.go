package main

import (
	"context"
	"log"
	"os"

	"github.com/carloscontrerasruiz/gorest/models"
	"github.com/carloscontrerasruiz/gorest/repository"
	"github.com/carloscontrerasruiz/gorest/routes"
	"github.com/carloscontrerasruiz/gorest/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var repo repository.CrudRepository[models.User]

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Env vars was not loaded")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal("New server was not created")
	}

	repo = repository.NewUserRepository()

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", routes.HomeHandler(s)).Methods("GET")
	r.HandleFunc("/signup", routes.SignUpHander(s, repo)).Methods("POST")
}
