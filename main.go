package main

import (
	"fmt"
	"net/http"
)


func main() {
	settings := config.NewSettings()
	http.HandleFunc("/", ServeRandomEmoji)
	http.ListenAndServe(":80", nil)
}

func ServeRandomEmoji(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "random emoji")
}
