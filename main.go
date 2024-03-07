package main

import (
	"flag"
)

func main(){
	listenAddr :=flag.String("listenAddr", ":3000", "listen address running on this port")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)

	server.Run()

	// price, err := svc.FetchPrice(context.Background(),"ETH")
 
	// if err != nil {
	// 	slog.Error(err.Error())
	// }
	// fmt.Println(price)
}