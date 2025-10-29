package repository

import "github.com/bruhngl0/goBruh/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
