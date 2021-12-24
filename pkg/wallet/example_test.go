package wallet

import (
	"reflect"
	"testing"
	"github.com/Munirkhuja/bank/v2/pkg/types"
)

func TestService_FindAccountByID(t *testing.T) {
	accounts:=[]types.Account{
		{ID: 1, Balance: 5_00,Phone: "992928882211"},
		{ID: 2, Balance: 50_00,Phone: "992928882212"},
		{ID: 3, Balance: 10_00,Phone: "992928882213"},
		{ID: 4, Balance: 16_00,Phone: "992928882214"},
	}
	service:=Service{nextAccountId: 5,accounts: accounts}
	expected := &types.Account{ID: 3, Balance: 10_00,Phone: "992928882213"}
	result,err := service.FindAccountByID(5)
	if err==ErrAccountNotFound {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v actual: %v",expected,result)
	}
}