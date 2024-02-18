package postgres

import (
	"github.com/OurLuv/l0/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderStorage interface {
	Create(model.Order) error
	GetAll() []model.Order
}

type OrderRepository struct {
	pool *pgxpool.Pool
}

func (o *OrderRepository) Create(model.Order) error {
	return nil
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		pool: pool,
	}
}
