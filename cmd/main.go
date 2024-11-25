package main

import (
	"fmt"

	models "github.com/TDanilV16/test-data-generator/models"
)

func main() {
	rnd := *models.RandomOrder()
	fmt.Printf("id: %s, amount: %d, date: %s, customerId: %s", rnd.Id.String(), rnd.PurchaseAmount, rnd.Date.String(), rnd.CustomerId.String())
}
