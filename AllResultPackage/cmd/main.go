package main

import (
	"log"
	"net/http"

	"github.com/mserebryaakov/course-golang/build/server"
)

func main() {
	http.HandleFunc("/", server.Handler)

	srv := new(server.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error when starting the server: %s", err.Error())
	}
}
