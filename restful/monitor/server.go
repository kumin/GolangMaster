package monitor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kumin/GolangMaster/restful/configs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

type MetricServer struct {
	port int
}

func NewMetricServer(
	config *configs.ServerConfiguration,
) *MetricServer {
	return &MetricServer{
		port: config.MetricPort,
	}
}

func (m *MetricServer) Start(ctx context.Context) error {
	log.Info().Msgf("start metric server on port: %d", m.port)
	serv := &http.Server{
		Addr: fmt.Sprintf(":%d", m.port),
	}
	http.Handle("/metric", promhttp.Handler())
	errCh := make(chan error, 1)
	defer close(errCh)
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			errCh <- err
		}
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
