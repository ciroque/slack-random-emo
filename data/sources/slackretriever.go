package sources

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"slack-random-emo/config"
	"slack-random-emo/data"
	"time"
)

type SlackRetriever struct {
	EmoUpdateChannel chan<- *[]data.Emo
}

func (retriever *SlackRetriever) Run(settings *config.Settings, stopChannel <-chan bool) {
	periodic := time.NewTicker(time.Second * settings.RetrievalPeriodSeconds)

	retriever.worker(settings)

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
					retriever.worker(settings)
				}
			}
		}
	}()
}

func (retriever *SlackRetriever) worker(settings *config.Settings) {
	url := fmt.Sprintf("%s?token=%s", settings.SlackUrl, settings.SlackAuthToken)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Error(err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	var content SlackEmoji
	err = json.Unmarshal(body, &content)
	if err != nil {
		logrus.Error(err)
		return
	}

	var emojis []data.Emo
	for name := range content.Emoji {
		emojis = append(emojis, data.Emo{name})
	}

	retriever.EmoUpdateChannel <- &emojis
}

type SlackEmoji struct {
	Ok       bool
	Emoji    map[string]string
	Cache_ts string
}
