package main

import (
	"context"
	"strings"

	"github.com/farischt/micro/types"
)

type PriceService interface {
	GetPrice(context.Context, string) (float64, error)
}

func NewPriceService() PriceService {
	return &priceService{}
}

type priceService struct {
}

func (s *priceService) parseCoin(coin string) string {
	return strings.ToUpper(coin)
}

func (s *priceService) GetPrice(ctx context.Context, coin string) (float64, error) {
	price, exist := priceDatabaseMock[s.parseCoin(coin)]

	if !exist {
		return price, types.NewUnsupportedCoinError(coin, "unsupported coin")
	}

	return price, nil
}
