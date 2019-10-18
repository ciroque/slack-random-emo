package main

import (
	"github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"slack-random-emo/data"
	"slack-random-emo/http"
	"syscall"
)

func main() {
	errorChannel := make(chan string)
	defer close(errorChannel)

	server := http.Server{
		Logger: logrus.NewEntry(logrus.New()),
		Emos:   []data.Emo{{Name: "one"}, {Name: "two"}, {Name: "three"}, {Name: "four"}, {Name: "five"}},
	}

	go server.Run(errorChannel)

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)

	select {
	case <-sigTerm:
		{

		}
	case err := <-errorChannel:
		{
			logrus.Error(err)
		}
	}
}
