package service

import (
	"fmt"
	"github.com/SubochevaValeriya/microservice-balance"
	"github.com/SubochevaValeriya/microservice-balance/pkg/repository"
	"github.com/mattevans/dinero"
	"os"
	"time"
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

	} else {
		client := dinero.NewClient(
			os.Getenv("OPEN_EXCHANGE_APP_ID"),
			"RUB",
			20*time.Minute,
		)

		rsp, err := client.Rates.Get(ccy)
		if err != nil {
			return balance, fmt.Errorf("can't convert balance to inputted CCY: %w", err)
		}
		fmt.Println(rsp)
	}
	//url := "https://api.apilayer.com/exchangerates_data/convert?to={to}&from={from}&amount={amount}"
	//
	//client := &http.Client{}
	//req, err := http.NewRequest("GET", url, nil)
	//req.Header.Set("apikey", "4EIJgz5GX9n5N8QUdRXHwQE01DfqGSqs")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//res, err := client.Do(req)
	//if res.Body != nil {
	//	defer res.Body.Close()
	//}
	//body, err := ioutil.ReadAll(res.Body)
	//
	//fmt.Println(string(body))
	//
	//json.Unmarshal(responseData, &responseObject)

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
