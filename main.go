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
	abortChannel := make(chan string)
	defer close(abortChannel)

	server := http.Server{
		Logger: logrus.NewEntry(logrus.New()),
		Emos:   []data.Emo{{Name: "one"}, {Name: "two"}, {Name: "three"}, {Name: "four"}, {Name: "five"}},
	}

	go server.Run(abortChannel)

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)

	select {
	case <-sigTerm:
		{

		}
	case err := <-abortChannel:
		{
			logrus.Error(err)
		}
	}
}
