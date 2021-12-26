package types

type Money int64
type Currency string
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
)
type PAN string
type Card struct {
	ID uint
	PAN PAN
	Balance Money
	MinBalance Money
	Currency Currency
	Color string
	Name string
	Active bool
}
type PaymentCategory string
type PaymentStatus string
const (
	PaymentStatusOk PaymentStatus="OK"
	PaymentStatusFail PaymentStatus="FAIL"
	PaymentStatusInProgress PaymentStatus="INPROGRESS"
)
type Payment struct {
	ID string
	AccountID int64
	Amount Money
	Category PaymentCategory
	Status PaymentStatus
}
type PaymentSource struct {
	Type string // 'card'
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}
type Phone string
type Account struct {
	ID int64
	Phone Phone
	Balance Money
}
type Favorite struct{
	ID string
	AccountID int64
	Name string
	Amount Money
	Category PaymentCategory
}