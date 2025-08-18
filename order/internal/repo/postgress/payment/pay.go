package payment

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"order/internal/apperr"
	"order/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
)

type PaymentRepo struct {
	DB *sqlx.DB
}

func NewPaymentRepo(db *sqlx.DB) *PaymentRepo {
	return &PaymentRepo{
		DB: db,
	}
}

type dto struct {
	ID           string    `db:"id"`
	OrderUID     string    `db:"order_uid"`
	Transaction  string    `db:"transaction"`
	RequestID    string    `db:"request_id"`
	Currency     string    `db:"currency"`
	Provider     string    `db:"provider"`
	Amount       int       `db:"amount"`
	PaymentDT    time.Time `db:"payment_dt"`
	Bank         string    `db:"bank"`
	DeliveryCost int       `db:"delivery_cost"`
	GoodsTotal   int       `db:"goods_total"`
	CustomFee    int       `db:"custom_fee"`
}

func (p *PaymentRepo) CreatePayment(ctx context.Context, payment model.Payment) (model.Payment, error) {
	query := `INSERT INTO payments (
		id, order_uid, transaction, request_id, currency, provider,
		amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
	)
	ON CONFLICT (order_uid) DO UPDATE SET
		transaction = EXCLUDED.transaction,
		request_id = EXCLUDED.request_id,
		currency = EXCLUDED.currency,
		provider = EXCLUDED.provider,
		amount = EXCLUDED.amount,
		payment_dt = EXCLUDED.payment_dt,
		bank = EXCLUDED.bank,
		delivery_cost = EXCLUDED.delivery_cost,
		goods_total = EXCLUDED.goods_total,
		custom_fee = EXCLUDED.custom_fee
	RETURNING id, order_uid, transaction, request_id, currency, provider,
	          amount, payment_dt, bank, delivery_cost, goods_total, custom_fee;`

	var tmp dto

	err := p.DB.GetContext(ctx, &tmp, query,
		payment.ID,
		payment.OrderUID,
		payment.Transaction,
		payment.RequestID,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDT,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	)
	if err != nil {
		return model.Payment{}, err
	}

	return model.Payment{
		ID:           tmp.ID,
		OrderUID:     tmp.OrderUID,
		Transaction:  tmp.Transaction,
		RequestID:    tmp.RequestID,
		Currency:     tmp.Currency,
		Provider:     tmp.Provider,
		Amount:       tmp.Amount,
		PaymentDT:    tmp.PaymentDT,
		Bank:         tmp.Bank,
		DeliveryCost: tmp.DeliveryCost,
		GoodsTotal:   tmp.GoodsTotal,
		CustomFee:    tmp.CustomFee,
	}, nil
}

func (p *PaymentRepo) GetPaymentsByOrderUID(ctx context.Context, orderUID string) (model.Payment, error) {
	query := `SELECT 
		id, order_uid, transaction, request_id, currency, provider,
		amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
	FROM payments
	WHERE order_uid = $1`

	var tmp dto

	err := p.DB.GetContext(ctx, &tmp, query, orderUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Payment{}, apperr.ErrNotFound
		}
		return model.Payment{}, err
	}
	log.Println("repo payment", tmp)
	return model.Payment{
		ID:           tmp.ID,
		OrderUID:     tmp.OrderUID,
		Transaction:  tmp.Transaction,
		RequestID:    tmp.RequestID,
		Currency:     tmp.Currency,
		Provider:     tmp.Provider,
		Amount:       tmp.Amount,
		PaymentDT:    tmp.PaymentDT,
		Bank:         tmp.Bank,
		DeliveryCost: tmp.DeliveryCost,
		GoodsTotal:   tmp.GoodsTotal,
		CustomFee:    tmp.CustomFee,
	}, nil
}
