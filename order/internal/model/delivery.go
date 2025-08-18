package model

type Delivery struct {
	ID       string // UUID, можно не использовать в логике
	OrderUID string
	Name     string
	Phone    string
	ZIP      string
	City     string
	Address  string
	Region   string
	Email    string
}
