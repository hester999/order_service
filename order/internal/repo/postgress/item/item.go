package item

import (
	"context"
	"log"
	"order/internal/apperr"
	"order/internal/model"

	"github.com/jmoiron/sqlx"
)

type ItemRepo struct {
	db *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *ItemRepo {
	return &ItemRepo{
		db: db,
	}
}

type dto struct {
	ID         string `db:"id"`
	OrderID    string `db:"order_uid"`
	ChrtID     int    `db:"chrt_id"`
	TrackID    string `db:"track_number"`
	Price      int    `db:"price"`
	Rid        string `db:"rid"`
	Name       string `db:"name"`
	Sale       int    `db:"sale"`
	Size       string `db:"size"`
	TotalPrice int    `db:"total_price"`
	NmID       int    `db:"nm_id"`
	Brands     string `db:"brand"`
	Status     int    `db:"status"`
}

func (i *ItemRepo) CreateItem(ctx context.Context, item model.Item) (model.Item, error) {
	query := `INSERT INTO items (
		id, order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
	)
	ON CONFLICT (order_uid, chrt_id) DO UPDATE SET
		track_number = EXCLUDED.track_number,
		price = EXCLUDED.price,
		rid = EXCLUDED.rid,
		name = EXCLUDED.name,
		sale = EXCLUDED.sale,
		size = EXCLUDED.size,
		total_price = EXCLUDED.total_price,
		nm_id = EXCLUDED.nm_id,
		brand = EXCLUDED.brand,
		status = EXCLUDED.status
	RETURNING id, order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status;`

	var tmp dto

	err := i.db.GetContext(ctx, &tmp, query,
		item.ID,
		item.OrderUID,
		item.ChrtID,
		item.TrackNumber,
		item.Price,
		item.RID,
		item.Name,
		item.Sale,
		item.Size,
		item.TotalPrice,
		item.NmID,
		item.Brand,
		item.Status)

	if err != nil {
		log.Println(err)
		return model.Item{}, err
	}

	return model.Item{
		ID:          tmp.ID,
		OrderUID:    tmp.OrderID,
		ChrtID:      tmp.ChrtID,
		TrackNumber: tmp.TrackID,
		Price:       tmp.Price,
		RID:         tmp.Rid,
		Name:        tmp.Name,
		Sale:        tmp.Sale,
		Size:        tmp.Size,
		TotalPrice:  tmp.TotalPrice,
		NmID:        tmp.NmID,
		Brand:       tmp.Brands,
		Status:      tmp.Status,
	}, nil
}

func (i *ItemRepo) GetItemsByOrderUID(ctx context.Context, orderUID string) ([]model.Item, error) {
	query := `SELECT 
		id, order_uid, chrt_id, track_number, price, rid, name, sale, size, 
		total_price, nm_id, brand, status 
	FROM items 
	WHERE order_uid = $1`

	var tmp []dto

	err := i.db.SelectContext(ctx, &tmp, query, orderUID)
	if err != nil {
		log.Printf("GetItemsByOrderUID failed: %v", err)
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, apperr.ErrNotFound
	}

	items := make([]model.Item, 0, len(tmp))
	for _, r := range tmp {
		items = append(items, model.Item{
			ID:          r.ID,
			OrderUID:    r.OrderID,
			ChrtID:      r.ChrtID,
			TrackNumber: r.TrackID,
			Price:       r.Price,
			RID:         r.Rid,
			Name:        r.Name,
			Sale:        r.Sale,
			Size:        r.Size,
			TotalPrice:  r.TotalPrice,
			NmID:        r.NmID,
			Brand:       r.Brands,
			Status:      r.Status,
		})
	}
	log.Println("repo item", items)
	return items, nil
}
