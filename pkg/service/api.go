package service

import (
	"database/sql"
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

func (s *ApiService) GetAllUsersBalances(user microservice.UsersBalances) (*sql.Row, error) {

	return s.repo.GetAllUsersBalances(user)
}

func (s *ApiService) GetBalanceById(user microservice.UsersBalances) (*sql.Row, error) {

	return s.repo.GetAllUsersBalances(user)
}
