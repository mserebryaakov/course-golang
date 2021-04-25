package main

import (
	"homework"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	srv := new(homework.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error when starting the server: %s", err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("url " + r.URL.Path + " client " + r.RemoteAddr)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.Write(w)
}
