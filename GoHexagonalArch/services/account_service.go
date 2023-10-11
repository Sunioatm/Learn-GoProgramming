package services

import (
	"goarch/errs"
	"goarch/logs"
	"goarch/repositories"
	"strings"
	"time"
)

type accountService struct {
	accountRepo repositories.AccountRepository
}

func NewAccountService(accountRepo repositories.AccountRepository) AccountService {
	return accountService{accountRepo: accountRepo}
}

func (s accountService) NewAccount(customerID int, req NewAccountRequest) (*AccountResponse, error) {
	// Validate
	if req.Amount < 5000 {
		return nil, errs.NewValidationError("amount should be greater or equal to 5000")
	}
	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be checking or saving")
	}

	account := repositories.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      1,
	}

	newAccount, err := s.accountRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	resAcc := AccountResponse{
		AccountID:   newAccount.AccountID,
		OpeningDate: newAccount.OpeningDate,
		AccountType: newAccount.AccountType,
		Amount:      newAccount.Amount,
		Status:      newAccount.Status,
	}
	return &resAcc, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accountRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	res := []AccountResponse{}
	for _, a := range accounts {
		res = append(res, AccountResponse{
			AccountID:   a.AccountID,
			OpeningDate: a.OpeningDate,
			AccountType: a.AccountType,
			Amount:      a.Amount,
			Status:      a.Status,
		})
	}

	return res, nil
}
