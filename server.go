package main

import (
	"context"
	"fmt"
	"net"

	"github.com/farischt/micro/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func startGRPC(s PriceService, addr uint, done chan<- bool) error {
	priceServiceServer := NewPriceServiceServer(s)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", addr))

	if err != nil {
		logrus.Error(err)
		return err
	}

	opts := []grpc.ServerOption{}
	logrus.WithFields(logrus.Fields{
		"port": addr,
	}).Info("GRPC Server starting:")
	server := grpc.NewServer(opts...)
	proto.RegisterPriceServiceServer(server, priceServiceServer)

	close(done)
	return server.Serve(ln)
}

type PriceServiceServer struct {
	service PriceService
	proto.UnimplementedPriceServiceServer
}

func NewPriceServiceServer(service PriceService) *PriceServiceServer {
	return &PriceServiceServer{
		service,
		proto.UnimplementedPriceServiceServer{},
	}
}

func (s *PriceServiceServer) GetPrice(ctx context.Context, r *proto.PriceRequest) (*proto.PriceResponse, error) {
	price, err := s.service.GetPrice(ctx, r.Coin)

	if err != nil {
		return nil, err
	}

	response := new(proto.PriceResponse)
	response.Coin = r.Coin
	response.Price = float32(price)

	return response, nil
}
