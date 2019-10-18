package config

import (
	"fmt"
	"os"
	"strconv"
)

type Settings struct {
	Host string
	Port int
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

	config := &Settings{host, nport}

	return config, nil
}
