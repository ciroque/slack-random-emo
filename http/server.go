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

func (server *Server) Run(stopCh <-chan struct{}) {
	settings, _ := config.NewSettings()
	server.Logger.Info("Starting worker..")
	server.worker(settings)
	<-stopCh
}

func (server *Server) worker(settings *config.Settings) {
	server.Logger.Info("About to start listening...")
	http.HandleFunc("/", server.ServeRandomEmoji)
	address := fmt.Sprintf("%s:%d", settings.Host, settings.Port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		server.Logger.Error(err)
	}
}

func (server *Server) ServeRandomEmoji(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "random emoji")
}
