package main

import (
	"flag"
)

func main() {
	port := flag.Uint("port", 3000, "The listening port")
	flag.Parse()

	service := NewLoggingService(NewPriceService())
	server := NewJsonApi(service, *port)
	server.Start()

}
