package model

import (
	"errors"
	"time"
)

type MarketList struct {
	Products []Product
	Created  time.Time
	Date     time.Time
	Market   string
}

func NewMarketList(market string, date time.Time, products []Product) (*MarketList, error) {

	if market == "" {
		return nil, errors.New("Market is mandatory")
	}

	return &MarketList{
		Products: products,
		Created:  time.Now(),
		Date:     date,
		Market:   market,
	}, nil
}
