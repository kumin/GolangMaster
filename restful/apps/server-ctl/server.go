package apps

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/kumin/GolangMaster/restful/configs"
	"github.com/kumin/GolangMaster/restful/handler"
)

type HttpServer struct {
	port   int
	logger *log.Logger
}

func NewHttpServer(
	configs *configs.ServerConfiguration,
	prodHandler *handler.ProductCtlHandler,
) *HttpServer {
	server := &HttpServer{
		port:   configs.Port,
		logger: log.Default(),
	}
	server.RegisterHandler("/v1/product/add", handler.HandlerWrapper(prodHandler.AddProduct))
	server.RegisterHandler("/v1/product/listing", handler.HandlerWrapper(prodHandler.ListProducts))
	server.RegisterHandler("/v1/product", handler.HandlerWrapper(prodHandler.GetProduct))
	return server
}

func (h *HttpServer) Start(ctx context.Context) {
	h.logger.Printf("Server is listening on port:%d", h.port)
	http.ListenAndServe(fmt.Sprintf(":%d", h.port), nil)
}

func (h *HttpServer) RegisterHandler(path string, handler http.HandlerFunc) {
	http.Handle(path, handler)
}
