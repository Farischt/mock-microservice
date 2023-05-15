package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceService
}

func NewLoggingService(next PriceService) PriceService {
	return &loggingService{
		next,
	}
}

func (s *loggingService) GetPrice(ctx context.Context, coin string) (price float64, err error) {
	defer func(begin time.Time) {
		log.WithFields(log.Fields{
			"took":  time.Since(begin),
			"coin":  coin,
			"price": price,
			"error": err,
		}).Info("Get price")
	}(time.Now())
	return s.next.GetPrice(ctx, coin)
}
