package postgres

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.com/OurLuv/l0/internal/model"
	"github.com/OurLuv/l0/internal/pubsub"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	pool, err = NewPool("postgres://postgres:admin@localhost:5432/wbl0")
	if err != nil {
		log.Printf("failed to init storage: %s", err)
		os.Exit(1)
	}
	//defer pool.Close()
	log.Printf("Storage init")
	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestOrderCreate(t *testing.T) {
	or := NewOrderRepository(pool)
	order := pubsub.RandomOrder()
	var o *model.Order
	var err error
	if o, err = or.Create(order); err != nil {
		_ = o
		t.Errorf("Error: %s", err)
	}
}
