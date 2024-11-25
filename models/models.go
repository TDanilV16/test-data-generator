package models

import (
	"github.com/TDanilV16/test-data-generator/pkg/date"
	"github.com/TDanilV16/test-data-generator/pkg/random"
	"github.com/google/uuid"
)

type Order struct {
	Id             uuid.UUID
	PurchaseAmount int
	Date           *date.Date
	CustomerId     uuid.UUID
}

func RandomOrder() *Order {
	id := uuid.New()
	amount := random.Amount(250, 10000)
	date := date.RandomDate()
	customerId := uuid.New()

	return &Order{Id: id, PurchaseAmount: amount, Date: date, CustomerId: customerId}
}
