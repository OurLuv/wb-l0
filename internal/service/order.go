package service

import (
	"github.com/OurLuv/l0/internal/model"
	"github.com/OurLuv/l0/internal/storage/cache"
	"github.com/OurLuv/l0/internal/storage/postgres"
	"github.com/google/uuid"
)

type OrderServcie interface {
	CreateInDataBase(order model.Order) error
	CreateInCache(order model.Order) error
	Save(order model.Order) error
	GetById(id uuid.UUID) (*model.Order, error)
	GetByIdCache(id uuid.UUID) (*model.Order, error)
}

type Order struct {
	repo  postgres.OrderStorage
	cache cache.OrderCache
}

func (o *Order) Save(order model.Order) error {
	// saving in database
	err := o.repo.Create(order)
	if err != nil {
		return err
	}

	//saving in cache
	o.cache.Put(order)

	return nil
}

func New(repo postgres.OrderStorage) *Order {
	return &Order{
		repo: repo,
	}
}
