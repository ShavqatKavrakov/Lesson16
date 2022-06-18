package types

//Money представляет собой денежную единицу в минимальных единицах (копейки,дирамы и.т)
type Money int64

//PaymentCategory представляет собой информацию о категориях который совершён платёж (автобб аптеки, рестороны и.т)
type PaymentCategory string

//PaymentStatus представляет статус платёжа
type PaymentStatus string

//статус платёжа
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "Fail"
	PaymentStatusInProgress PaymentStatus = "InProgress"
)

//Payment представдят информацию о платёже
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

type Phone string

//Account представлет информацию о счёте ползователья
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

//Favorite представляет платёж в избранное.
type Favorite struct {
	ID        string
	AccountID int64
	Name      string
	Amount    Money
	Category  PaymentCategory
}