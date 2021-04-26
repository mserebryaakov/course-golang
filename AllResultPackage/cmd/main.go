package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mserebryaakov/course-golang/build/server"
)

func main() {
	http.HandleFunc("/", handler)

	srv := new(server.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error when starting the server: %s", err.Error())
	}
}

//Функция, обрабатывающая запрос "/"
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("Error read body: %s", err.Error())
			return
		}
		w.Write(body)
	}
}
