package main

import (
	"github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"slack-random-emo/http"
	"syscall"
)

func main() {
	stopCh := make(chan struct{})
	defer close(stopCh)

	server := http.Server{Logger: logrus.NewEntry(logrus.New())}

	go server.Run(stopCh)

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	<-sigTerm
}
