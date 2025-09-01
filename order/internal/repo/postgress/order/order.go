package order

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

type dto struct {
	OrderUID          string    `db:"order_uid"`
	TrackNumber       string    `db:"track_number"`
	Entry             string    `db:"entry"`
	Locale            string    `db:"locale"`
	InternalSignature string    `db:"internal_signature"`
	CustomerID        string    `db:"customer_id"`
	DeliveryService   string    `db:"delivery_service"`
	ShardKey          string    `db:"shard_key"`
	SmID              int       `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OOFShard          string    `db:"oof_shard"`
}
type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (o *OrderRepo) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {

	query := `INSERT INTO orders (
		order_uid, track_number, entry, locale, internal_signature, 
		customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	ON CONFLICT (order_uid) DO UPDATE SET
		track_number = EXCLUDED.track_number,
		entry = EXCLUDED.entry,
		locale = EXCLUDED.locale,
		internal_signature = EXCLUDED.internal_signature,
		customer_id = EXCLUDED.customer_id,
		delivery_service = EXCLUDED.delivery_service,
		shard_key = EXCLUDED.shard_key,
		sm_id = EXCLUDED.sm_id,
		date_created = EXCLUDED.date_created,
		oof_shard = EXCLUDED.oof_shard
	RETURNING order_uid, track_number, entry, locale, internal_signature, 
	          customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard;`

	var tmp dto

	err := o.db.GetContext(ctx, &tmp, query,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.ShardKey,
		order.SMID,
		order.DateCreated,
		order.OOFShard,
	)
	if err != nil {
		log.Println(err)
		return model.Order{}, err
	}

	result := model.Order{
		OrderUID:          tmp.OrderUID,
		TrackNumber:       tmp.TrackNumber,
		Entry:             tmp.Entry,
		Locale:            tmp.Locale,
		InternalSignature: tmp.InternalSignature,
		CustomerID:        tmp.CustomerID,
		DeliveryService:   tmp.DeliveryService,
		ShardKey:          tmp.ShardKey,
		SMID:              tmp.SmID,
		DateCreated:       tmp.DateCreated,
		OOFShard:          tmp.OOFShard,
	}
	return result, nil
}

func (o *OrderRepo) GetOrderByOrderUID(ctx context.Context, id string) (model.Order, error) {

	query := `SELECT order_uid, track_number, entry, locale, internal_signature, 
		customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard FROM orders WHERE order_uid = $1`

	var tmp dto

	err := o.db.GetContext(ctx, &tmp, query, id)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return model.Order{}, apperr.ErrNotFound
		}
		return model.Order{}, err
	}

	result := model.Order{
		OrderUID:          tmp.OrderUID,
		TrackNumber:       tmp.TrackNumber,
		Entry:             tmp.Entry,
		Locale:            tmp.Locale,
		InternalSignature: tmp.InternalSignature,
		CustomerID:        tmp.CustomerID,
		DeliveryService:   tmp.DeliveryService,
		ShardKey:          tmp.ShardKey,
		SMID:              tmp.SmID,
		DateCreated:       tmp.DateCreated,
		OOFShard:          tmp.OOFShard,
	}
	return result, nil
}

func (o *OrderRepo) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	query := `SELECT order_uid, track_number, entry, locale, internal_signature, customer_id,
		delivery_service, shard_key, sm_id, date_created, oof_shard FROM orders`

	var tmp []dto
	err := o.db.SelectContext(ctx, &tmp, query)
	if err != nil {
		log.Println(err)
		return []model.Order{}, err
	}

	if len(tmp) == 0 {
		return []model.Order{}, apperr.ErrNotFound
	}

	orders := make([]model.Order, len(tmp))

	for i, item := range tmp {
		orders[i] = model.Order{
			OrderUID:          item.OrderUID,
			TrackNumber:       item.TrackNumber,
			Entry:             item.Entry,
			Locale:            item.Locale,
			InternalSignature: item.InternalSignature,
			CustomerID:        item.CustomerID,
			DeliveryService:   item.DeliveryService,
			ShardKey:          item.ShardKey,
			SMID:              item.SmID,
			DateCreated:       item.DateCreated,
			OOFShard:          item.OOFShard,
		}
	}
	return orders, nil
}
