package service

import (
	"fmt"
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

func (s *ApiService) GetBalanceById(userId int, ccy string) (microservice.UsersBalances, error) {
	balance, err := s.repo.GetBalanceById(userId)
	if err != nil {
		return balance, err
	}

	if ccy != "" {
		balance.Balance, err = convertToCCY(ccy, balance.Balance)
		if err != nil {
			return balance, fmt.Errorf("can't convert balance to inputted CCY: %w", err)
		}
	}

	return balance, err
}

func (s *ApiService) DeleteUserById(userId int) error {

	return s.repo.DeleteUserById(userId)
}

func (s *ApiService) DeleteAllUsersBalances() error {

	return s.repo.DeleteAllUsersBalances()
}

func (s *ApiService) ChangeBalanceById(userId int, transaction microservice.Transactions) (microservice.Transactions, error) {

	return s.repo.ChangeBalanceById(userId, transaction)
}
