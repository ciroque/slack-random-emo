package config

import (
	"fmt"
	"os"
	"strconv"
)

type Settings struct {
	Port int
}

func NewSettings() (*Settings, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return &Settings{80}, nil
	}

	nport, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("unable to parse PORT: %v", err)
	}

	config := &Settings{nport}

	return config, nil
}
