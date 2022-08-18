package service

import "github.com/SubochevaValeriya/microservice-balance/pkg/repository"

type Balance interface {
}

type Service struct {
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
