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
	go startGRPC(service, *port, doneChan)
	go func(done <-chan bool) {
		logrus.Info("Waiting for GRPC server to start")
		<-done
		ctx := context.Background()
		grpcClient.GetPrice(ctx, &proto.PriceRequest{
			Coin: "ETH",
		})
	}(doneChan)

	// jsonClient := client.New(fmt.Sprintf("http://localhost:%d", 8000))

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ctx := context.Background()
	// 	jsonClient.GetCoinPrice(ctx, "ETH")
	// }()

	// server := NewJsonApi(service, 8000)
	// server.Start()
}
