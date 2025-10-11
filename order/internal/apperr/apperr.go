package apperr

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrInvalidUUID        = errors.New("uuid invalid")
	ErrUUIDIsEmpty        = errors.New("empty uuid")
	ErrTrackNumberIsEmpty = errors.New("track number is empty")
	ErrEntryIsEmpty       = errors.New("entry is empty")
	ErrLocationIsEmpty    = errors.New("location is empty")
	ErrCustomerIDIsEmpty  = errors.New("customer id is empty")
)

var (
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrNameIsEmpty        = errors.New("name is empty")
	ErrZIPIsEmpty         = errors.New("zip is empty")
	ErrAddressIsEmpty     = errors.New("address is empty")
	ErrCityIsEmpty        = errors.New("city is empty")
	ErrInvalidEmail       = errors.New("invalid email")
)

var (
	ErrInvalidDelivery        = errors.New("delivery is invalid or missing")
	ErrInvalidPayment         = errors.New("payment is invalid or missing")
	ErrEmptyItems             = errors.New("items list cannot be empty")
	ErrEmailIsEmpty           = errors.New("email is empty")
	ErrDeliveryServiceIsEmpty = errors.New("delivery service is empty")
)

var (
	ErrTransactionIsEmpty  = errors.New("transaction is empty")
	ErrCurrencyIsEmpty     = errors.New("currency is empty")
	ErrProviderIsEmpty     = errors.New("provider is empty")
	ErrInvalidAmount       = errors.New("invalid amount")
	ErrInvalidDeliveryCost = errors.New("invalid delivery cost")
	ErrInvalidGoodsTotal   = errors.New("invalid goods total")
)

var (
	ErrInvalidChrtID    = errors.New("invalid chrt_id")
	ErrItemNameIsEmpty  = errors.New("item name is empty")
	ErrInvalidItemPrice = errors.New("invalid item price")
)
