package validator

import (
	"net/mail"
	"order/internal/apperr"
	"order/internal/consumer/dto"

	"github.com/google/uuid"
	"github.com/nyaruka/phonenumbers"
)

// ValidateIncomingOrder — общая валидация всего входящего заказа
func ValidateIncomingOrder(order dto.OrderDTO) error {

	if err := validateOrder(order); err != nil {
		return err
	}

	// Проверяем доставку
	if order.Delivery == nil {
		return apperr.ErrInvalidDelivery
	}
	if err := validateDelivery(*order.Delivery); err != nil {
		return err
	}

	// Проверяем оплату
	if order.Payment == nil {
		return apperr.ErrInvalidPayment
	}
	if err := validatePayment(*order.Payment); err != nil {
		return err
	}

	if len(order.Items) == 0 {
		return apperr.ErrEmptyItems
	}
	for _, item := range order.Items {
		if err := validateItem(item); err != nil {
			return err
		}
	}

	return nil
}

func validateOrder(order dto.OrderDTO) error {
	if order.OrderUID == "" {
		return apperr.ErrUUIDIsEmpty
	}

	if err := uuid.Validate(order.OrderUID); err != nil {
		return apperr.ErrInvalidUUID
	}

	if order.TrackNumber == "" {
		return apperr.ErrTrackNumberIsEmpty
	}
	if order.Entry == "" {
		return apperr.ErrEntryIsEmpty
	}
	if order.Locale == "" {
		return apperr.ErrLocationIsEmpty
	}
	if order.CustomerID == "" {
		return apperr.ErrCustomerIDIsEmpty
	}
	if order.DeliveryService == "" {
		return apperr.ErrDeliveryServiceIsEmpty
	}
	return nil
}

func validateDelivery(delivery dto.DeliveryDTO) error {
	if delivery.Name == "" {
		return apperr.ErrNameIsEmpty
	}

	phone, err := phonenumbers.Parse(delivery.Phone, "RU")
	if err != nil || !phonenumbers.IsValidNumberForRegion(phone, "RU") {
		return apperr.ErrInvalidPhoneNumber
	}

	if delivery.ZIP == "" {
		return apperr.ErrZIPIsEmpty
	}
	if delivery.Address == "" {
		return apperr.ErrAddressIsEmpty
	}
	if delivery.City == "" {
		return apperr.ErrCityIsEmpty
	}
	if delivery.Email == "" {
		return apperr.ErrEmailIsEmpty
	}

	if _, err = mail.ParseAddress(delivery.Email); err != nil {
		return apperr.ErrInvalidEmail
	}

	return nil
}

func validatePayment(payment dto.PaymentDTO) error {
	if payment.Transaction == "" {
		return apperr.ErrTransactionIsEmpty
	}
	if payment.Currency == "" {
		return apperr.ErrCurrencyIsEmpty
	}
	if payment.Provider == "" {
		return apperr.ErrProviderIsEmpty
	}
	if payment.Amount <= 0 {
		return apperr.ErrInvalidAmount
	}
	if payment.DeliveryCost < 0 {
		return apperr.ErrInvalidDeliveryCost
	}
	if payment.GoodsTotal < 0 {
		return apperr.ErrInvalidGoodsTotal
	}
	return nil
}

func validateItem(item dto.ItemDTO) error {
	if item.ChrtID == 0 {
		return apperr.ErrInvalidChrtID
	}
	if item.TrackNumber == "" {
		return apperr.ErrTrackNumberIsEmpty
	}
	if item.Name == "" {
		return apperr.ErrItemNameIsEmpty
	}
	if item.Price < 0 {
		return apperr.ErrInvalidItemPrice
	}
	return nil
}
