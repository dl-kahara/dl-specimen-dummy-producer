package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	MetricsNamespace = "datalounges"
	MetricsSubsystem = "producer"
)

var (
	requests_produced *prometheus.CounterVec
	bytes_produced    *prometheus.CounterVec
)

func SetupMetrics() {
	requests_produced = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: MetricsNamespace,
		Subsystem: MetricsSubsystem,
		Name:      "requests_total",
	}, []string{})

	bytes_produced = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: MetricsNamespace,
		Subsystem: MetricsSubsystem,
		Name:      "bytes_total",
	}, []string{})
}

func Metrics(address string) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal().Err(err).Str("address", address).Msg("Could not expose Prometheus metrics")
	}
}
