package model

import "time"

//type Payment struct {
//	ID           string
//	Transaction  string
//	RequestId    string
//	Currency     string
//	Provider     string
//	Amount       int
//	PaymentDt    time.Time
//	Bank         string
//	DeliveryCost int
//	GoodsTotal   int
//	CustomFee    int
//}

type Payment struct {
	ID           string
	OrderUID     string
	Transaction  string
	RequestID    string
	Currency     string
	Provider     string
	Amount       int
	PaymentDT    time.Time
	Bank         string
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}
