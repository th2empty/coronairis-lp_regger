package service

import (
	"coronairis-lp_web-app/configs"
	"coronairis-lp_web-app/pkg/repository"
)

var (
	_ = configs.InitConfig()
)

type RegistrationService struct {
	repo repository.Registration
}

func NewRegistrationService(repo repository.Registration) *RegistrationService {
	return &RegistrationService{repo: repo}
}

func (s *RegistrationService) RegisterUser(token string, uid int) error {
	return s.repo.RegisterUser(token, uid)
}

func (s *RegistrationService) UpdateUser(token string, uid int) error {
	return s.repo.UpdateUser(token, uid)
}
