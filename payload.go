package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"time"
)

type Payload struct {
	Timestamp string `json:"timestamp"`
	Prime     string `json:"prime"`
}

func GeneratePayload(primeBits int) []byte {
	prime, err := rand.Prime(rand.Reader, primeBits)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate a random prime, giving up")
	}
	payload := Payload{
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Prime:     hex.EncodeToString(prime.Bytes()),
	}
	marshalled, err := json.Marshal(payload)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to marshall payload, giving up")
	}

	return marshalled
}
