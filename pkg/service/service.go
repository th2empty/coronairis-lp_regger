package service

import (
	"coronairis-lp_web-app/pkg/repository"
)

type Registration interface {
	RegisterUser(token string, uid int) error
	UpdateUser(token string, uid int) error
}

type Service struct {
	Registration
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Registration: NewRegistrationService(repos.Registration),
	}
}
