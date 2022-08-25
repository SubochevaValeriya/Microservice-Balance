package service

import (
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/SubochevaValeriya/microservice-balance/pkg/repository"
)

type Balance interface {
	CreateUser(user microservice.UsersBalances) (int, error)
	GetAllUsersBalances() ([]microservice.UsersBalances, error)
	GetBalanceById(userId int) (microservice.UsersBalances, error)
}

type Service struct {
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{newApiService(repos.Balance)}
}
