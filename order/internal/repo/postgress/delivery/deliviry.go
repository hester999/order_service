package delivery

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"order/internal/apperr"
	"order/internal/model"

	"github.com/jmoiron/sqlx"
)

type DeliveryRepo struct {
	db *sqlx.DB
}

func NewDeliveryRepo(db *sqlx.DB) *DeliveryRepo {
	return &DeliveryRepo{db: db}
}

type dto struct {
	ID       string `db:"id"`
	OrderUID string `db:"order_uid"`
	Name     string `db:"name"`
	Phone    string `db:"phone"`
	ZIP      string `db:"zip"`
	City     string `db:"city"`
	Address  string `db:"address"`
	Region   string `db:"region"`
	Email    string `db:"email"`
}

func (d *DeliveryRepo) CreateDelivery(ctx context.Context, delivery model.Delivery) (model.Delivery, error) {
	query := `INSERT INTO deliveries (id,order_uid, name,phone,zip,city,address,region,email) 
			  VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
			  ON CONFLICT (order_uid) DO UPDATE SET
			  name = EXCLUDED.name,
			  phone = EXCLUDED.phone,
			  zip = EXCLUDED.zip,
			  city = EXCLUDED.city,
			  address = EXCLUDED.address,
			  region = EXCLUDED.region,
			  email = EXCLUDED.email
			  RETURNING id,order_uid, name,phone,zip,city,address,region,email `

	var tmp dto
	err := d.db.GetContext(ctx, &tmp, query, delivery.ID, delivery.OrderUID, delivery.Name, delivery.Phone, delivery.ZIP, delivery.City, delivery.Address, delivery.Region, delivery.Email)

	if err != nil {
		log.Println(err)
		return model.Delivery{}, err
	}

	return model.Delivery{
		ID:       tmp.ID,
		OrderUID: tmp.OrderUID,
		Name:     tmp.Name,
		Phone:    tmp.Phone,
		ZIP:      tmp.ZIP,
		City:     tmp.City,
		Address:  tmp.Address,
		Region:   tmp.Region,
		Email:    tmp.Email,
	}, nil

}

func (d *DeliveryRepo) GetDeliveryByOrderUID(ctx context.Context, orderUID string) (model.Delivery, error) {
	var tmp dto

	query := `SELECT id,order_uid, name,phone,zip,city,address,region,email FROM deliveries WHERE order_uid = $1`

	err := d.db.GetContext(ctx, &tmp, query, orderUID)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return model.Delivery{}, apperr.ErrNotFound
		}
		return model.Delivery{}, err
	}

	log.Println("repo delivery", tmp)
	return model.Delivery{
		ID:       tmp.ID,
		OrderUID: tmp.OrderUID,
		Name:     tmp.Name,
		Phone:    tmp.Phone,
		ZIP:      tmp.ZIP,
		City:     tmp.City,
		Address:  tmp.Address,
		Region:   tmp.Region,
		Email:    tmp.Email,
	}, nil

}
