package http

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"net/http"
	"slack-random-emo/config"
)

type Server struct {
	Logger *logrus.Entry
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
	fmt.Fprintf(writer, "random emoji")
}
