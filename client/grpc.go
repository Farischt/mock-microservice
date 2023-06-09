package client

import (
	"context"

	"github.com/farischt/micro/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCService(addr string) (proto.PriceServiceClient, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	client := proto.NewPriceServiceClient(conn)

	return client, nil
}

type GRPCService struct {
	client proto.PriceServiceClient
}

func NewGRPCServiceClient(client proto.PriceServiceClient) *GRPCService {
	return &GRPCService{
		client,
	}
}

func (c *GRPCService) GetPrice(ctx context.Context, coin string) (*proto.PriceResponse, error) {
	return c.client.GetPrice(ctx, &proto.PriceRequest{
		Coin: coin,
	})
}

func (c *GRPCService) RemoveCoin(ctx context.Context, coin string) (*proto.RemoveCoinResponse, error) {
	return c.client.RemoveCoin(ctx, &proto.RemoveCoinRequest{
		Coin: coin,
	})
}
