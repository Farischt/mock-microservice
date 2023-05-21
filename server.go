package main

import (
	"context"
	"fmt"
	"net"

	"github.com/farischt/micro/proto"
	"github.com/farischt/micro/types"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
	price, err := s.service.GetPrice(ctx, r.GetCoin())

	if err != nil {
		if _, ok := err.(types.UnsupportedCoinError); ok {
			return nil, grpc.Errorf(codes.NotFound, err.Error())
		}

		return nil, grpc.Errorf(codes.Internal, "unknown error")
	}

	response := new(proto.PriceResponse)
	response.Coin = r.GetCoin()
	response.Price = float32(price)

	return response, nil
}

func (s *PriceServiceServer) RemoveCoin(ctx context.Context, r *proto.RemoveCoinRequest) (*proto.RemoveCoinResponse, error) {
	err := s.service.RemoveCoin(ctx, r.GetCoin())

	if err != nil {
		if _, ok := err.(types.UnsupportedCoinError); ok {
			return nil, grpc.Errorf(codes.NotFound, err.Error())
		}

		return nil, grpc.Errorf(codes.Internal, "unknown error")
	}

	response := new(proto.RemoveCoinResponse)
	response.Coin = r.GetCoin()
	response.Success = true

	return response, nil

}
