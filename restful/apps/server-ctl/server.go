package apps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kumin/GolangMaster/restful/configs"
	"github.com/kumin/GolangMaster/restful/handler"
	"github.com/rs/zerolog/log"
)

type HttpServer struct {
	port int
}

func NewHttpServer(
	configs *configs.ServerConfiguration,
	prodHandler *handler.ProductCtlHandler,
) *HttpServer {
	server := &HttpServer{
		port: configs.APIPort,
	}
	server.RegisterHandler("/v1/product/add", handler.HandlerWrapper(prodHandler.AddProduct))
	server.RegisterHandler("/v1/product/listing", handler.HandlerWrapper(prodHandler.ListProducts))
	server.RegisterHandler("/v1/product", handler.HandlerWrapper(prodHandler.GetProduct))
	return server
}

func (h *HttpServer) Start(ctx context.Context) error {
	log.Info().Msgf("start HTTP server on port: %d", h.port)
	serv := &http.Server{Addr: fmt.Sprintf(":%d", h.port)}
	errCh := make(chan error, 1)
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			errCh <- err
		}
		close(errCh)
	}()
	for {
		select {
		case <-ctx.Done():
			return serv.Shutdown(ctx)
		case err := <-errCh:
			return err
		}
	}
}

func (h *HttpServer) RegisterHandler(path string, handler http.HandlerFunc) {
	http.Handle(path, handler)
}
