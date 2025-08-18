package model

import "time"

type Order struct {
	OrderUID          string
	TrackNumber       string
	Entry             string
	Delivery          *Delivery
	Payment           *Payment
	Items             []Item
	Locale            string
	InternalSignature string
	CustomerID        string
	DeliveryService   string
	ShardKey          string
	SMID              int
	DateCreated       time.Time
	OOFShard          string
}
