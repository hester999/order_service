package order

import "time"

type ItemDTO struct {
	ID          string `json:"id,omitempty"`
	OrderUID    string `json:"order_uid,omitempty"`
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type DeliveryDTO struct {
	ID       string `json:"id,omitempty"`
	OrderUID string `json:"order_uid,omitempty"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	ZIP      string `json:"zip"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Region   string `json:"region"`
	Email    string `json:"email"`
}

type PaymentDTO struct {
	ID           string `json:"id,omitempty"`
	OrderUID     string `json:"order_uid,omitempty"`
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type OrderDTO struct {
	OrderUID          string       `json:"order_uid"`
	TrackNumber       string       `json:"track_number"`
	Entry             string       `json:"entry"`
	Delivery          *DeliveryDTO `json:"delivery"`
	Payment           *PaymentDTO  `json:"payment"`
	Items             []ItemDTO    `json:"items"`
	Locale            string       `json:"locale"`
	InternalSignature string       `json:"internal_signature"`
	CustomerID        string       `json:"customer_id"`
	DeliveryService   string       `json:"delivery_service"`
	ShardKey          string       `json:"shardkey"`
	SMID              int          `json:"sm_id"`
	DateCreated       time.Time    `json:"date_created"`
	OOFShard          string       `json:"oof_shard"`
}
