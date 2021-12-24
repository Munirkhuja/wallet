package wallet

import (
	"errors"
	"github.com/Munirkhuja/bank/v2/pkg/types"
)
var ErrAccountNotFound=errors.New("account not found")
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