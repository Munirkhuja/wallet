package wallet

import (
	"reflect"
	"testing"

	"github.com/Munirkhuja/wallet/pkg/types"
	"github.com/google/uuid"
)
func TestService_Reject(t *testing.T) {
	uuidTest:=uuid.New().String()
	accounts:=[]*types.Account{
		{ID: 1, Balance: 5_00,Phone: "992928882211"},
		{ID: 2, Balance: 50_00,Phone: "992928882212"},
		{ID: 3, Balance: 10_00,Phone: "992928882213"},
		{ID: 4, Balance: 16_00,Phone: "992928882214"},
	}
	payments:=[]*types.Payment{
		{ID: uuidTest,AccountID: 2,Amount: 30_00, Status: types.PaymentStatusInProgress},
		{ID: uuid.New().String(),AccountID: 2,Amount: 3_00, Status: types.PaymentStatusInProgress},
		{ID: uuid.New().String(),AccountID: 1,Amount: 10_00, Status: types.PaymentStatusOk},
		{ID: uuid.New().String(),AccountID: 2,Amount: 30_00, Status: types.PaymentStatusInProgress},
		{ID: uuid.New().String(),AccountID: 4,Amount: 15_00, Status: types.PaymentStatusInProgress},
		{ID: uuid.New().String(),AccountID: 2,Amount: 90_00, Status: types.PaymentStatusFail},
		{ID: uuid.New().String(),AccountID: 1,Amount: 1_000_00, Status: types.PaymentStatusOk},
	}	
	service:=Service{nextAccountId: 5,accounts: accounts,payments: payments}
	expected := &types.Payment{ID: uuidTest,AccountID: 2,Amount: 30_00, Status: types.PaymentStatusFail}
	result,err := service.Reject(uuidTest)
	if err==ErrAccountNotFound {
		t.Error(err)
		return
	}
	if err==ErrPaymentNotFound {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v actual: %v",expected,result)
	}
}
func TestService_FindAccountByID(t *testing.T) {
	accounts:=[]*types.Account{
		{ID: 1, Balance: 5_00,Phone: "992928882211"},
		{ID: 2, Balance: 50_00,Phone: "992928882212"},
		{ID: 3, Balance: 10_00,Phone: "992928882213"},
		{ID: 4, Balance: 16_00,Phone: "992928882214"},
	}
	service:=Service{nextAccountId: 5,accounts: accounts}
	expected := &types.Account{ID: 3, Balance: 10_00,Phone: "992928882213"}
	result,err := service.FindAccountByID(3)
	if err==ErrAccountNotFound {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v actual: %v",expected,result)
	}
}