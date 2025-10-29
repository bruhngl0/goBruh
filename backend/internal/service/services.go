package service

import (
	"github.com/bruhngl0/goBruh/internal/lib/job"
	"github.com/bruhngl0/goBruh/internal/repository"
	"github.com/bruhngl0/goBruh/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {

	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
