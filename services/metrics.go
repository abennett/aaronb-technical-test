package services

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func ExposeMetrics(l *zap.Logger) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(":8411", mux); err != nil {
			l.Error("metrics server error", zap.Error(err))
		}
	}()
}

var (
	requestsReceived = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_total",
		},
		[]string{"method"},
	)
	responsesSent = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "responses_total",
		}, []string{"status_code"},
	)
)
