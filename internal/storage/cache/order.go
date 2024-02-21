package cache

import (
	"fmt"
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

func (oc *OrderCache) Get(uuid uuid.UUID) (*model.Order, error) {
	v, ok := oc.store.Load(uuid)
	res, ok1 := v.(model.Order)
	if !ok1 {
		return nil, fmt.Errorf("Error in cache")
	}
	if !ok {
		return nil, fmt.Errorf("This is no order with this uuid")
	}
	return &res, nil
}

func (oc *OrderCache) Put(order model.Order) {
	oc.store.Store(order.OrderUUID, order)
}

func (oc *OrderCache) Recover([]model.Order) {
	// todo finish it
}

func New() *OrderCache {
	return &OrderCache{}
}
