package main

import (
	"context"
	"fmt"
	"log/slog"
)

func main(){

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	price, err := svc.FetchPrice(context.Background(),"ETH")
 
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Println(price)
}