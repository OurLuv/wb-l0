package service

import (
	"github.com/OurLuv/l0/internal/model"
	"github.com/OurLuv/l0/internal/storage/cache"
	"github.com/OurLuv/l0/internal/storage/postgres"
	"github.com/google/uuid"
)

type OrderServcie interface {
	Save(model.Order) (string, error)
	GetById(string) (*model.Order, error)
	Pull() error
}

type Order struct {
	repo  postgres.OrderStorage
	cache cache.OrderCacheInterface
}

func (o *Order) Save(order model.Order) (string, error) {
	// saving in database
	orderFullInfo, err := o.repo.Create(order)
	if err != nil {
		return "", err
	}

	//saving in cache
	o.cache.Put(*orderFullInfo)

	return orderFullInfo.OrderUUID.String(), nil
}

func (o *Order) GetById(uuidStr string) (*model.Order, error) {
	//parsing
	uuid, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, err
	}

	//getting order from cache
	order, err := o.cache.Get(uuid)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Pull() error {
	orders, err := o.repo.GetAll()
	if err != nil {
		return err
	}

	o.cache.Recover(orders)

	return nil
}

func New(repo postgres.OrderStorage, cache cache.OrderCacheInterface) *Order {
	return &Order{
		repo:  repo,
		cache: cache,
	}
}
