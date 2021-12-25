package wallet

import (
	"errors"

	"github.com/Munirkhuja/wallet/pkg/types"
)
var ErrAccountNotFound=errors.New("account not found")
var ErrPhoneRegistered=errors.New("phone already registered")
var ErrAmountMustBePositive= errors.New("amount must be greater then zero")
type Service struct{
	nextAccountId int64
	accounts []types.Account
	payments []types.Payment
}
func (s *Service) FindAccountByID(accountID int64) (*types.Account, error)  {
	for _, account :=range s.accounts{
		if account.ID==accountID {
			return &account,nil
		}
	}
	return nil,ErrAccountNotFound
}
func (s *Service) RegisterAccount(phone types.Phone)(*types.Account,error){
	for _, account :=range s.accounts{
		if phone==account.Phone {
			return nil,ErrPhoneRegistered
		}
	}
	s.nextAccountId++
	account:=&types.Account{
		ID: s.nextAccountId,
		Phone: phone,
		Balance: 0,
	}
	s.accounts=append(s.accounts, *account)
	return account,nil
}