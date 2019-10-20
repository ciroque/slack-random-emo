package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Settings struct {
	Host                   string
	Port                   int
	RetrievalPeriodSeconds time.Duration
}

func NewSettings() (*Settings, error) {
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "888"
	}

	nport, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("unable to parse PORT: %v", err)
	}

	retrievalPeriod := os.Getenv("RETRIEVAL_PERIOD_SECONDS")
	if retrievalPeriod == "" {
		retrievalPeriod = "60"
	}

	nRetrievalPeriod, err := strconv.Atoi(retrievalPeriod)
	if err != nil {
		return nil, fmt.Errorf("unable to parse RETRIEVAL_PERIOD_SECONDS: %v", err)
	}

	config := &Settings{host, nport, time.Duration(nRetrievalPeriod)}

	return config, nil
}
