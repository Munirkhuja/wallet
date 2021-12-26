package wallet

import (
	"errors"

	"github.com/Munirkhuja/wallet/pkg/types"
	"github.com/google/uuid"
)
var ErrAccountNotFound=errors.New("account not found")
var ErrPhoneRegistered=errors.New("phone already registered")
var ErrAmountMustBePositive= errors.New("amount must be greater then zero")
var ErrFavoriteNotFound=errors.New("favorite not found")
var ErrPaymentNotFound=errors.New("payment not found")
var ErrPaymentAlreadyRejected=errors.New("payment already rejected")
var ErrNotEnoughBalance=errors.New("not enough balance")
type Service struct{
	nextAccountId int64
	accounts []*types.Account
	payments []*types.Payment
	favorite []*types.Favorite
}
func (s *Service) FindAccountByID(accountID int64) (*types.Account, error)  {
	for _, account :=range s.accounts{
		if account.ID==accountID {
			return account,nil			
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
	s.accounts=append(s.accounts, account)
	return account,nil
}
func (s *Service) Deposit(accountID int64,amount types.Money)error{
	if amount<=0 {
		return ErrAmountMustBePositive
	}
	account,_:=s.FindAccountByID(accountID)
	if account==nil {
		return ErrAccountNotFound
	}
	account.Balance+=amount
	return nil
}

func(s *Service)Pay(accountID int64, amount types.Money, category types.PaymentCategory)(*types.Payment,error){
	if amount<=0 {
		return nil,ErrAmountMustBePositive
	}
	account,_:=s.FindAccountByID(accountID)
	if account==nil {
		return nil,ErrAccountNotFound
	}
	if account.Balance<amount {
		return nil,ErrNotEnoughBalance
	}
	account.Balance-=amount
	paymentID:=uuid.New().String()
	payment:=&types.Payment{
		ID: paymentID,
		AccountID: accountID,
		Amount: amount,
		Category: category,
		Status: types.PaymentStatusInProgress,
	}
	s.payments=append(s.payments, payment)
	return payment,nil
}
func (s *Service)Reject(paymentID string)  error {
	payment,err:=s.FindPaymentByID(paymentID)
	if err==nil {		
		if payment.Status==types.PaymentStatusFail {
			return 	ErrPaymentAlreadyRejected	
		}
		account,_:=s.FindAccountByID(payment.AccountID)
		if account==nil {
			return ErrAccountNotFound
		}
		account.Balance+=payment.Amount
		payment.Status=types.PaymentStatusFail
	}
	return err
}
func (s *Service)FindPaymentByID(paymentID string) (*types.Payment, error) {
	for _, payment :=range s.payments{
		if paymentID==payment.ID {
			return payment,nil			
		}
	}
	return nil,ErrPaymentNotFound
}
func(s *Service)Repeat(paymentID string)(*types.Payment, error){
	payment,_:=s.FindPaymentByID(paymentID)
	if payment==nil {
		return nil,ErrPaymentNotFound
	}
	newPayment,err:=s.Pay(payment.AccountID,payment.Amount,payment.Category)
	return newPayment,err
}
func (s *Service)FavoritePayment(paymentID string, name string) (*types.Favorite, error){
	payment,_:=s.FindPaymentByID(paymentID)
	if payment==nil {
		return nil,ErrPaymentNotFound
	}
	favoriteID:=uuid.New().String()
	favorite:=&types.Favorite{
		ID: favoriteID,
		Name: "OLE",
		AccountID: payment.AccountID,
		Amount: payment.Amount,
		Category: payment.Category,
	}
	s.favorite=append(s.favorite,favorite )
	return favorite,nil
}
func(s *Service)PayFromFavorite(favoriteID string) (*types.Payment, error){
	favorite,_:=s.FindFavoriteByID(favoriteID)
	if favorite==nil {
		return nil,ErrFavoriteNotFound
	}
	payment,err:=s.Pay(favorite.AccountID,favorite.Amount,favorite.Category)
	return payment,err
}
func (s *Service)FindFavoriteByID(favoriteID string) (*types.Favorite, error) {
	for _, favorite :=range s.favorite{
		if favoriteID==favorite.ID {
			return favorite,nil			
		}
	}
	return nil,ErrFavoriteNotFound
}