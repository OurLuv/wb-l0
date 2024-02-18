package pubsub

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/OurLuv/l0/internal/model"
	uuid "github.com/google/uuid"
)

func randomOrder() model.Order {
	tn := "WBILM" + randomString(7)
	order := model.Order{
		OrderUUID:         uuid.UUID{},
		TrackNumber:       tn,
		Entry:             "WBIL",
		Delivery:          randomDelivery(),
		Payment:           randomPayment(),
		Items:             randomItems(tn),
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        randomString(3),
		DeliveryService:   "meest",
		ShardKey:          randomString(1),
		SmID:              randomInt(1, 100),
		DateCreated:       time.Now(),
		OofShard:          randomString(1),
	}
	return order
}

func randomPayment() model.Payment {
	payment := model.Payment{
		Transaction:  randomUIID(),
		RequestID:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       randomInt(1, 5000),
		PaymentDate:  time.Now().Unix(),
		Bank:         "alpha",
		DeliveryCost: randomInt(200, 1500),
		GoodsTotal:   randomInt(200, 1500),
		CustomFee:    0,
	}

	return payment
}

func randomDelivery() model.Delivery {
	delivery := model.Delivery{
		Id:      randomInt(1, 20),
		Name:    randomString(10),
		Phone:   randomPhone(),
		Zip:     randomString(7),
		City:    randomString(8),
		Address: randomString(10),
		Region:  randomString(10),
		Email:   randomEmail(),
	}

	return delivery
}

func randomItems(trackNumber string) []model.Item {
	items := []model.Item{}
	for i := 0; i < randomInt(1, 5); i++ {
		item := model.Item{
			ChrtId:      randomInt(100, 10000),
			TrackNumber: trackNumber,
			Price:       randomInt(1000, 5000),
			RID:         randomString(20),
			Name:        randomString(10),
			Sale:        randomInt(0, 51),
			Size:        randomString(3),
			TotalPrice:  randomInt(250, 1600),
			NmID:        randomInt(100000, 1500000),
			Brand:       randomString(10),
			Status:      randomInt(100, 200),
		}
		items = append(items, item)
	}

	return items
}

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789"

func randomString(l int) string {
	a := make([]byte, l)
	for i := range a {
		a[i] = charset[rand.Intn(len(charset))]
	}
	return string(a)
}

func randomEmail() string {
	email := randomString(6)
	email += "@gmail.com"

	return email
}

func randomPhone() string {
	p := "+"
	for i := 0; i < 11; i++ {
		p += strconv.Itoa(randomInt(0, 9))
	}
	return p
}

func randomUIID() uuid.UUID {
	return uuid.New()
}
