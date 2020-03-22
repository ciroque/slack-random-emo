package http

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"slack-random-emo/config"
	"slack-random-emo/data"
)

type Server struct {
	AbortChannel     chan<- string
	Logger           *logrus.Entry
	Emos             *[]data.Emo
	EmoUpdateChannel <-chan *[]data.Emo
	Settings         *config.Settings
}

func (server *Server) Run() {
	http.HandleFunc("/", server.ServeRandomEmoji)
	http.Handle("/metrics", promhttp.Handler())
	address := fmt.Sprintf("%s:%d", server.Settings.Host, server.Settings.Port)
	server.Logger.Info("Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		server.AbortChannel <- err.Error()
	}
}

func (server *Server) ServeRandomEmoji(writer http.ResponseWriter, request *http.Request) {
	length := len(*server.Emos)
	index := rand.Intn(length)

	response := SlackEmojiResponse{
		ResponseType: "in_channel",
		Text:         fmt.Sprintf(":%s:", (*server.Emos)[index].Name),
		Attachments:  []map[string]string{},
	}

	bytes, err := json.Marshal(&response)
	if err != nil {
		server.Logger.Warnf("Error responding to request %#v", err)
	}

	writer.Header().Add("Content-Type", "application/json")
	_, err = fmt.Fprintf(writer, "%s", bytes)
	if err != nil {
		server.Logger.Warnf("Error responding to request %#v", err)
	}
}

func (server *Server) HandleUpdates() {
	for updatedEmos := range server.EmoUpdateChannel {
		server.Emos = updatedEmos
	}
}
