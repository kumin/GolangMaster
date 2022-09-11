package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	port   int
	logger *log.Logger
}

func NewHttpServer(port int) *HttpServer {
	return &HttpServer{
		port:   port,
		logger: log.Default(),
	}
}

func (h *HttpServer) Start(ctx context.Context) {
	h.logger.Printf("Server is listening on port:%d", h.port)
	http.ListenAndServe(fmt.Sprintf(":%d", h.port), nil)
}

func (h *HttpServer) RegisterHandler(path string, handler http.HandlerFunc) {
	http.Handle(path, handler)
}
