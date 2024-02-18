package cache

import (
	"sync"

	"github.com/OurLuv/l0/internal/model"
	"github.com/google/uuid"
)

type OrderCacheInterface interface {
	Get(uuid.UUID) (*model.Order, error)
	Put(model.Order)
	Recover([]model.Order)
}

type OrderCache struct {
	store sync.Map
}

func (oc *OrderCache) Get(uuid.UUID) (*model.Order, error) {
	return nil, nil
}

func (oc *OrderCache) Put(order model.Order) {
	oc.store.Store(order.OrderUUID, order)
}

func (oc *OrderCache) Recover([]model.Order) {
	// todo finish it
}
