package main

import (
	"github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"slack-random-emo.org/config"
	"slack-random-emo.org/data"
	"slack-random-emo.org/data/sources"
	"slack-random-emo.org/metrics"
	localHttp "slack-random-emo.org/server"
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

	metrics := metrics.NewMetrics()

	var emos *[]data.Emo

	server := localHttp.Server{
		AbortChannel:     abortChannel,
		Logger:           logrus.NewEntry(logrus.New()),
		Emos:             emos,
		EmoUpdateChannel: emoUpdateChannel,
		Settings:         settings,
		Metrics:          &metrics,
	}

	slackEmoRetriever := sources.SlackRetriever{
		EmoUpdateChannel: emoUpdateChannel,
		Settings:         settings,
		StopChannel:      stopRetrieverChannel,
		Metrics:          &metrics,
	}

	go server.Run()
	go server.HandleUpdates()
	go slackEmoRetriever.Run()

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
