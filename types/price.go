package types

import "time"

type PriceResponse struct {
	Coin  string  `json:"coin"`
	Price float64 `json:"price"`
}

type UnsupportedCoinError struct {
	Coin      string    `json:"coin"`
	Err       string    `json:"err"`
	Timestamp time.Time `json:"timestamp"`
}

func NewUnsupportedCoinError(coin string, e string) UnsupportedCoinError {
	return UnsupportedCoinError{
		Coin:      coin,
		Err:       e,
		Timestamp: time.Now(),
	}
}

func (a UnsupportedCoinError) Error() string {
	return a.Err
}
