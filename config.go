package main

import (
	"os"
	"strconv"
)

const (
	DefaultPrimeBits      = 4096
	DefaultConsumer       = "http://localhost:8080"
	DefaultMetricsAddress = ":9108"
)

type Config struct {
	PrimeBits      int
	Consumer       string
	MetricsAddress string
}

func NewConfig() *Config {
	var config Config

	primeBits := os.Getenv("PRIMEBITS")
	if primeBits == "" {
		config.PrimeBits = DefaultPrimeBits
	} else {
		bits, _ := strconv.Atoi(primeBits)
		config.PrimeBits = bits
	}

	consumer := os.Getenv("CONSUMER")
	if consumer == "" {
		config.Consumer = DefaultConsumer
	} else {
		config.Consumer = consumer
	}

	metricsAddress := os.Getenv("METRICS_ADDRESS")
	if metricsAddress == "" {
		config.MetricsAddress = DefaultMetricsAddress
	} else {
		config.MetricsAddress = metricsAddress
	}

	return &config
}
