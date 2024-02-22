package postgres

import (
	"context"

	"github.com/OurLuv/l0/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderStorage interface {
	Create(model.Order) (*model.Order, error)
	GetAll() ([]model.Order, error)
}

type OrderRepository struct {
	pool *pgxpool.Pool
}

func (or *OrderRepository) Create(order model.Order) (*model.Order, error) {
	tx, err := or.pool.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	// delivery
	query := "INSERT INTO delivery (name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	row := tx.QueryRow(context.Background(), query, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if row.Scan(&order.Delivery.Id); err != nil {
		return nil, err
	}

	//payment
	query = `
		INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`

	row = tx.QueryRow(context.Background(), query, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDate, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if row.Scan(&order.Payment.Id); err != nil {
		return nil, err
	}

	// order
	query = `
	INSERT INTO orders (track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING order_uuid`

	row = tx.QueryRow(context.Background(), query,
		order.TrackNumber, order.Entry, order.Delivery.Id, order.Payment.Id, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard)

	if row.Scan(&order.OrderUUID); err != nil {
		return nil, err
	}

	//items
	query = "INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	queryManyToMany := "INSERT INTO orders_items (order_uuid, item_id) VALUES ($1, $2)"
	for index, i := range order.Items {
		row := tx.QueryRow(context.Background(), query, i.ChrtId, i.TrackNumber, i.Price, i.RID, i.Name, i.Sale, i.Size, i.TotalPrice, i.NmID, i.Brand, i.Status)
		if err := row.Scan(&order.Items[index].Id); err != nil {
			return nil, err
		}
		if _, err = tx.Exec(context.Background(), queryManyToMany, order.OrderUUID, order.Items[index].Id); err != nil {
			return nil, err
		}
	}

	return &order, nil
}

// todo: finish item's sql
func (or *OrderRepository) GetAll() ([]model.Order, error) {
	var order model.Order
	orders := []model.Order{}

	tx, err := or.pool.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()
	query := "SELECT * from orders o join delivery d on o.delivery_id = d.id join payment p on o.payment_id = p.id"
	rows, err := tx.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&order.OrderUUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Delivery.Id,
			&order.Payment.Id,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
			&order.Delivery.Id,
			&order.Delivery.Name,
			&order.Delivery.Phone,
			&order.Delivery.Zip,
			&order.Delivery.City,
			&order.Delivery.Address,
			&order.Delivery.Region,
			&order.Delivery.Email,
			&order.Payment.Transaction,
			&order.Payment.RequestID,
			&order.Payment.Currency,
			&order.Payment.Provider,
			&order.Payment.Amount,
			&order.Payment.PaymentDate,
			&order.Payment.Bank,
			&order.Payment.DeliveryCost,
			&order.Payment.GoodsTotal,
			&order.Payment.CustomFee,
			&order.Payment.Id)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	//fetch items
	for i := range orders {
		items := []model.Item{}
		itemsQuery := "SELECT id, chrt_id, track_number, price, rid, i.name, sale, size, total_price, nm_id, brand, i.status from items i " +
			"LEFT JOIN orders_items oi ON i.id = oi.item_id " +
			"WHERE oi.order_uuid = $1"
		itemRows, err := tx.Query(context.Background(), itemsQuery, orders[i].OrderUUID)
		if err != nil {
			return nil, err
		}
		for itemRows.Next() {
			var item model.Item
			err := itemRows.Scan(
				&item.Id,
				&item.ChrtId,
				&item.TrackNumber,
				&item.Price,
				&item.RID,
				&item.Name,
				&item.Sale,
				&item.Size,
				&item.TotalPrice,
				&item.NmID,
				&item.Brand,
				&item.Status,
			)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
		orders[i].Items = items
	}
	return orders, nil
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		pool: pool,
	}
}
