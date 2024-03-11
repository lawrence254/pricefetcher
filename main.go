package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"github.com/lawrence254/pricefetcher/client"
)

func main(){
	client := client.New("http://localhost:3000")
	price, err := client.FetchPrice(context.Background(),"ET")

	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Printf("%+v\n", price)

	listenAddr :=flag.String("listenAddr", ":3000", "listen address running on this port")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)

	server.Run()


}