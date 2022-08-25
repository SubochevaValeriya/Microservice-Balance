package service

import (
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/SubochevaValeriya/microservice-balance/pkg/repository"
)

type ApiService struct {
	repo repository.Balance
}

func newApiService(repo repository.Balance) *ApiService {
	return &ApiService{repo: repo}
}

func (s *ApiService) CreateUser(user microservice.UsersBalances) (int, error) {

	return s.repo.CreateUser(user)
}

func (s *ApiService) GetAllUsersBalances() ([]microservice.UsersBalances, error) {

	return s.repo.GetAllUsersBalances()
}

func (s *ApiService) GetBalanceById(userId int) (microservice.UsersBalances, error) {

	return s.repo.GetBalanceById(userId)
}
