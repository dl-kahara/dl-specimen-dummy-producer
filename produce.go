package main

import (
	"bytes"
	"crypto/sha256"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"io"
	"net/http"
	"time"
)

func produce(config Config) {
	for {
		post(config)
	}
}

func post(config Config) []byte {
	var (
		err      error
		response *http.Response
		body     []byte
	)

	payload := GeneratePayload(config.PrimeBits)
	log.Debug().Bytes("payload", payload).Msg("POSTing payload")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if response, err = ctxhttp.Post(ctx, nil, config.Consumer, "application/json", bytes.NewReader(payload)); err != nil {
		log.Fatal().Err(err).Msg("Failed to POST payload")
	}
	defer response.Body.Close()

	if body, err = io.ReadAll(response.Body); err != nil {
		log.Fatal().Err(err).Msg("Failed to read response body")
	}

	bytes_produced.WithLabelValues().Add(float64(len(payload)))
	requests_produced.WithLabelValues().Inc()

	checksum := sha256.Sum256(payload)
	if bytes.Compare(body, checksum[:]) != 0 {
		log.Fatal().Str("status", response.Status).Any("headers", response.Header).Any("expected", checksum[:]).Any("got", body).Msg("Got an unexpected response from server, giving up")
	} else {
		log.Debug().Str("status", response.Status).Any("headers", response.Header).Any("expected", checksum[:]).Any("got", body).Msg("Got a valid response from server")
	}

	return body
}
