package http

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"math/rand"
	"net/http"
	"slack-random-emo/config"
	"slack-random-emo/data"
)

type Server struct {
	Logger           *logrus.Entry
	Emos             *[]data.Emo
	EmoUpdateChannel <-chan *[]data.Emo
}

func (server *Server) Run(settings *config.Settings, abortChannel chan<- string) {
	http.HandleFunc("/", server.ServeRandomEmoji)
	address := fmt.Sprintf("%s:%d", settings.Host, settings.Port)
	server.Logger.Info("Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		abortChannel <- err.Error()
	}
}

func (server *Server) ServeRandomEmoji(writer http.ResponseWriter, request *http.Request) {
	length := len(*server.Emos)
	index := rand.Intn(length)

	response := SlackEmojiResponse{
		ResponseType: "in_channel",
		Text:         (*server.Emos)[index].Name,
		Attachments:  []map[string]string{},
	}

	bytes, err := json.Marshal(&response)
	if err != nil {
		server.Logger.Warnf("Error responding to request %#v", err)
	}

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
