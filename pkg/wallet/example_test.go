package wallet

import (
	"reflect"
	"testing"

	"github.com/Munirkhuja/wallet/pkg/types"
	"github.com/google/uuid"
)
var testAccounts =[]*types.Account{
	{ID: 1, Balance: 5_00,Phone: "992928882211"},
	{ID: 2, Balance: 50_00,Phone: "992928882212"},
	{ID: 3, Balance: 10_00,Phone: "992928882213"},
	{ID: 4, Balance: 16_00,Phone: "992928882214"},
}
var testPayments=[]*types.Payment{
	{ID: uuid.New().String(),AccountID: 2,Amount: 30_00, Status: types.PaymentStatusInProgress},
	{ID: uuid.New().String(),AccountID: 6,Amount: 3_00, Status: types.PaymentStatusInProgress},
	{ID: uuid.New().String(),AccountID: 1,Amount: 10_00, Status: types.PaymentStatusOk},
	{ID: uuid.New().String(),AccountID: 2,Amount: 30_00, Status: types.PaymentStatusInProgress},
	{ID: uuid.New().String(),AccountID: 4,Amount: 15_00, Status: types.PaymentStatusInProgress},
	{ID: uuid.New().String(),AccountID: 2,Amount: 90_00, Status: types.PaymentStatusFail},
	{ID: uuid.New().String(),AccountID: 1,Amount: 1_000_00, Status: types.PaymentStatusOk},
}
var testServices=Service{nextAccountId: 5,accounts: testAccounts,payments: testPayments}
func TestService_Reject_success_loc(t *testing.T) {		
	s:=testServices
	payment:=testPayments[0]
	err := s.Reject(payment.ID)
	if err!=nil {
		t.Errorf("invalid result, %v",err)
	}
}
func TestService_Reject_account_not_found_loc(t *testing.T) {		
	s:=testServices	
	payment:=testPayments[1]
	err := s.Reject(payment.ID)
	if err!=ErrAccountNotFound {
		t.Errorf("invalid result, %v",err)
	}	
}
func TestService_Reject_already_rejected_loc(t *testing.T) {		
	s:=testServices
	payment:=testPayments[5]
	err := s.Reject(payment.ID)	
	if err!=ErrPaymentAlreadyRejected {
		t.Errorf("invalid result, %v",ErrPaymentAlreadyRejected)
	}
}
func TestService_FindAccountByID_success_loc(t *testing.T) {
	s:=testServices
	expected := &types.Account{ID: 3, Balance: 10_00,Phone: "992928882213"}
	result,_ := s.FindAccountByID(3)
	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v actual: %v",expected,result)
	}
}

func TestService_FindAccountByID_fail_loc(t *testing.T) {
	s:=testServices
	_,err := s.FindAccountByID(6)
	if err!=ErrAccountNotFound {
		t.Error(err)
	}
}
func TestService_FindPaymentByID_success_loc(t *testing.T) {
	s:=testServices	
	payment:=testPayments[2]
	expected := payment
	result,_ := s.FindPaymentByID(payment.ID)
	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v actual: %v",expected,result)
	}
}

func TestService_FindPaymentByID_fail_loc(t *testing.T) {
	s:=testServices
	_,err := s.FindPaymentByID(uuid.New().String())
	if err!=ErrPaymentNotFound {
		t.Error(err)
	}
}
func TestService_RegisterAccount_success_loc(t *testing.T) {
	s:=testServices
	phone:=types.Phone("992928882215")
	result,_:=s.RegisterAccount(phone)
	expected:=&types.Account{ID: s.nextAccountId,Phone: phone,Balance: 0}	
	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v actual: %v",expected,result)
	}
}
func TestService_RegisterAccount_fail_loc(t *testing.T) {
	s:=testServices
	phone:=types.Phone("992928882214")
	_,err:=s.RegisterAccount(phone)
	if err!=ErrPhoneRegistered {
		t.Errorf("invalid result, expected: %v actual: %v",ErrPhoneRegistered,err)
	}
}

func TestService_Deposit_success_loc(t *testing.T) {		
	s:=testServices
	account:=testAccounts[0]
	err := s.Deposit(account.ID,100_00)
	if err!=nil {
		t.Errorf("invalid result, %v",err)
	}
}
func TestService_Deposit_account_not_found_loc(t *testing.T) {		
	s:=testServices	
	err := s.Deposit(100,10_00)
	if err!=ErrAccountNotFound {
		t.Errorf("invalid result, %v",err)
	}	
}
func TestService_Deposit_fail_loc(t *testing.T) {		
	s:=testServices
	account:=testAccounts[0]
	err := s.Deposit(account.ID,-5_00)
	if err!=ErrAmountMustBePositive {
		t.Errorf("invalid result, %v",err)
	}
}
func TestService_Pay_success_loc(t *testing.T) {
	s:=testServices	
	account:=testAccounts[1]
	_,err := s.Pay(account.ID,50_00,"phone")
	if err!= nil {
		t.Errorf("invalid result,  %v",err)
	}
}
func TestService_Repeat_success_loc(t *testing.T) {
	s:=testServices	
	payment:=testPayments[4]
	_,err := s.Repeat(payment.ID)
	if err!= nil {
		t.Errorf("invalid result,  %v",err)
	}
}
func TestService_Repeat_account_not_found_loc(t *testing.T) {		
	s:=testServices	
	payment:=testPayments[1]
	_,err := s.Repeat(payment.ID)
	if err!=ErrAccountNotFound {
		t.Errorf("invalid result, %v",err)
	}	
}