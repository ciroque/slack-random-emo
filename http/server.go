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
	http.HandleFunc("/", server.ServeRandomEmoji)
	http.ListenAndServe(fmt.Sprintf("%s:%v", settings.Port), nil)
}

func (server *Server) ServeRandomEmoji(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "random emoji")
}
