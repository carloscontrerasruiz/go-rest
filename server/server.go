package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Brocker struct {
	config *Config
	router *mux.Router
}

func (b *Brocker) Config() *Config {
	return b.config
}

func (b *Brocker) Start(binder func(s Server, r *mux.Router)) {
	//b.router = mux.NewRouter()
	binder(b, b.router)
	log.Println("Starting new server", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("Server fail")
	}
}

func NewServer(ctx context.Context, config *Config) (*Brocker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("JWT secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("DB url is required")
	}

	return &Brocker{
		config: config,
		router: mux.NewRouter(),
	}, nil
}
