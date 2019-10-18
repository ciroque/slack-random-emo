package http

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"math/rand"
	"net/http"
	"slack-random-emo/config"
	"slack-random-emo/data"
)

type Server struct {
	Logger *logrus.Entry
	Emos   []data.Emo
}

func (server *Server) Run(errorChannel chan<- string) {
	settings, _ := config.NewSettings()
	http.HandleFunc("/", server.ServeRandomEmoji)
	address := fmt.Sprintf("%s:%d", settings.Host, settings.Port)
	server.Logger.Info("Listening on ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		errorChannel <- err.Error()
	}
}

func (server *Server) ServeRandomEmoji(writer http.ResponseWriter, request *http.Request) {
	length := len(server.Emos)
	index := rand.Intn(length)
	_, err := fmt.Fprintf(writer, "random emoji %d", index)
	if err != nil {
		server.Logger.Warnf("Error responding to request %#v", err)
	}
}
