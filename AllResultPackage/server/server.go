package server

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

//Функция, создающая слушающий http сервер
func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
	}
	return s.httpServer.ListenAndServe()
}

//Функция завершения работы сервера
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

//Функция, обрабатывающая запрос "/"
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("Error read body: %s", err.Error())
			return
		}
		w.Write(body)
	}
}
