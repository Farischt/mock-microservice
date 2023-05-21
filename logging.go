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

func getLogFields(begin time.Time, err error, baseFields log.Fields) log.Fields {
	fields := log.Fields{
		"took": time.Since(begin),
	}

	if err != nil {
		fields["error"] = err
	} else {
		for key, value := range baseFields {
			fields[key] = value
		}
	}

	return fields
}

func (s *loggingService) GetPrice(ctx context.Context, coin string) (price float64, err error) {
	defer func(begin time.Time) {
		fields := log.Fields{
			"coin": coin,
			"price": price,
		}
		log.WithFields(getLogFields(begin, err, fields)).Info("Get price")
	}(time.Now())
	return s.next.GetPrice(ctx, coin)
}

func (s *loggingService) RemoveCoin(ctx context.Context, coin string) (err error) {
	defer func(begin time.Time) {
		fields := log.Fields{
			"coin":  coin,
			"deleted": err == nil,
		}
		log.WithFields(getLogFields(begin, err, fields)).Info("Remove coin")
	}(time.Now())

	return s.next.RemoveCoin(ctx, coin)
}
