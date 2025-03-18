package model

import (
	"time"
)

type MarketList struct {
	Products []Product
	Created  time.Time
	Date     time.Time
	Market   string
}
