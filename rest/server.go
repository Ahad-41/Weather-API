package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"weather-api/config"
	weatherHandler "weather-api/rest/handlers/weather"
	"weather-api/rest/middlewares"
)

type Server struct {
	cnf            *config.Config
	weatherHandler *weatherHandler.Handler
}

func NewServer(
	cnf *config.Config,
	weatherHandler *weatherHandler.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		weatherHandler: weatherHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	
	// Register global middlewares
	manager.Use(
		middleware.Logger,
		middleware.RateLimitMiddleware(server.cnf),
	)

	mux := http.NewServeMux()
	
	// Register routes
	server.weatherHandler.RegisterRoutes(mux, manager)

	// Wrap mux with global middlewares
	wrappedMux := manager.WrapMux(mux)

	address := ":" + strconv.Itoa(server.cnf.HttpPort)

	fmt.Println("Weather API server running on", address)
	err := http.ListenAndServe(address, wrappedMux)
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
		os.Exit(1)
	}
}
