package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/farischt/micro/client"
	"github.com/farischt/micro/proto"
	"github.com/sirupsen/logrus"
)

func main() {
	port := flag.Uint("port", 3000, "The listening port")
	flag.Parse()

	service := NewLoggingService(NewPriceService())
	grpcClient, err := client.NewGRPCService(fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		panic(err)
	}

	doneChan := make(chan bool)
	errChan := make(chan error)
	go startGRPC(service, *port, doneChan, errChan) //nolint:errcheck

	go func() {
		select {
		case <-doneChan:
			logrus.Info("GRPC Server started")
			ctx := context.Background()
			_, _ = grpcClient.GetPrice(ctx, &proto.PriceRequest{
				Coin: "ETH",
			})
			_, _ = grpcClient.GetPrice(ctx, &proto.PriceRequest{
				Coin: "ET",
			})

		case err := <-errChan:
			logrus.Error(err)
			panic(err)
		}
	}()

	server := NewJsonApi(service, 8000)
	server.Start()
}
