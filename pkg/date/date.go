package date

import (
	"fmt"
	"strconv"
	"time"

	"github.com/TDanilV16/test-data-generator/pkg/random"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func RandomDate() *Date {
	randomMonth := random.Month()
	now := time.Now()
	randomYear := random.Year(now.Year())
	maxDay := time.Date(randomYear, randomMonth, 0, 0, 0, 0, 0, time.UTC).Day()
	randomDay := random.Day(maxDay)

	return &Date{
		Year: randomYear, Month: randomMonth, Day: randomDay,
	}
}

func (date Date) String() string {
	year := strconv.Itoa(date.Year)
	month := strconv.Itoa(int(date.Month))
	day := strconv.Itoa(date.Day)
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}
