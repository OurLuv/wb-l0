package service

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.com/OurLuv/l0/internal/pubsub"
	"github.com/OurLuv/l0/internal/storage/cache"
	"github.com/OurLuv/l0/internal/storage/postgres"
)

var s *Order

func TestMain(m *testing.M) {
	// storage
	var err error
	pool, err := postgres.NewPool("postgres://postgres:admin@localhost:5432/wbl0")
	if err != nil {
		log.Printf("failed to init storage: %s", err)
		os.Exit(1)
	}
	repo := postgres.NewOrderRepository(pool)
	//defer pool.Close()
	log.Printf("Storage init")

	//cache
	cache := cache.New()

	//service
	s = New(repo, cache)

	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestSaveAndGet(t *testing.T) {
	//generate order
	o := pubsub.RandomOrder()

	//saving
	uuidStr, err := s.Save(o)
	if err != nil {
		t.Error(err)
	}

	//getting
	res, err := s.GetById(uuidStr)
	if err != nil {
		t.Error(err)
	}
	_ = res
}
