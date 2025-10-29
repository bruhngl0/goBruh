package handler

import "github.com/bruhngl0/goBruh/internal/server"

type Handlers struct {
	server *server.Server
}

func NewHandlers(s *server.Server) *Handler {
	return &Handlers{
		server: s,
	}
}
