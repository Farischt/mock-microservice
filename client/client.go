package client

import (
	"context"

	"github.com/farischt/micro/types"
)

type IClient interface {
	GetCoinPrice(context.Context, string) (*types.ApiResponse[types.PriceResponse], error)
}
