package server

import (
	"context"
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
