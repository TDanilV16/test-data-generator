package random

import (
	"math/rand"
	"time"
)

func fromInterval(left, right int) int {
	return rand.Intn(right+1-left) + left
}

func Month() time.Month {
	return time.Month(fromInterval(1, 12))
}

func Year(nowYear int) int {
	return fromInterval(2018, nowYear)
}

func Day(maxDay int) int {
	return fromInterval(1, maxDay)
}

func Amount(minAmount, maxAmount int) int {
	return fromInterval(minAmount, maxAmount)
}
