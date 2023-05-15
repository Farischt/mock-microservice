package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/farischt/micro/types"
)

// Implements IClient
type Client struct {
	baseUrl string
}

func New(baseUrl string) IClient {
	return &Client{
		baseUrl,
	}
}

func (c *Client) GetCoinPrice(ctx context.Context, coin string) (*types.ApiResponse[types.PriceResponse], error) {
	endpoint := fmt.Sprintf("%s/coin/%s", c.baseUrl, coin)
	fmt.Println(endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	priceRes := new(types.ApiResponse[types.PriceResponse])
	err = json.NewDecoder(res.Body).Decode(priceRes)
	if err != nil {
		return nil, err
	}

	return priceRes, nil
}
