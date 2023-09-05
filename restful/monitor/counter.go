package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "request_counter",
		Help:        "this is just a test metric",
		ConstLabels: prometheus.Labels{"method": "get", "action": "listing"},
	})
	LatencyHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "latency_histogram",
		Help:    "this is just a test metric",
		Buckets: []float64{0.1, 0.5, 10.0, 15.0, 20.0},
	}, []string{"method", "action"})
)

func Init() {
	prometheus.MustRegister(LatencyHistogram)
}
