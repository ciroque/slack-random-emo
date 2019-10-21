package main

import (
	"github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"slack-random-emo/config"
	"slack-random-emo/data"
	"slack-random-emo/data/sources"
	"slack-random-emo/http"
	"syscall"
)

func main() {
	stopRetrieverChannel := make(chan bool)
	defer close(stopRetrieverChannel)

	abortChannel := make(chan string)
	defer close(abortChannel)

	emoUpdateChannel := make(chan *[]data.Emo)
	defer close(emoUpdateChannel)

	settings, _ := config.NewSettings()

	var emos *[]data.Emo

	server := http.Server{
		Logger:           logrus.NewEntry(logrus.New()),
		Emos:             emos,
		EmoUpdateChannel: emoUpdateChannel,
	}

	slackEmoRetriever := sources.SlackRetriever{
		EmoUpdateChannel: emoUpdateChannel,
	}

	go server.Run(settings, abortChannel)
	go server.HandleUpdates()
	go slackEmoRetriever.Run(settings, stopRetrieverChannel)

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)

	select {
	case <-sigTerm:
		{
			stopRetrieverChannel <- true
			logrus.Info("Exiting per SIGTERM")
		}
	case err := <-abortChannel:
		{
			stopRetrieverChannel <- true
			logrus.Error(err)
		}
	}
}
