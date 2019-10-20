package sources

import (
	"github.com/Sirupsen/logrus"
	"slack-random-emo/config"
	"time"
)

type SlackRetriever struct {
}

func (retriever *SlackRetriever) Run(settings *config.Settings, stopChannel <-chan bool) {
	periodic := time.NewTicker(time.Second * settings.RetrievalPeriodSeconds)

	go func() {
		for {
			select {
			case <-stopChannel:
				{
					logrus.Info("Shutting down Slack Retriever")
					return
				}
			case t := <-periodic.C:
				{
					logrus.Info("Tick at ", t)
					retriever.worker()
				}
			}
		}
	}()
}

func (retriever *SlackRetriever) worker() {
	logrus.Info("worker...")
}
