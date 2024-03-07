package main

import (
	"context"
	"fmt"
)
type PriceFetcher interface{
	FetchPrice(context.Context, string )(float64, error)
}

//Implements the PriceFetcher interface
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string)(float64, error){
	return MockPriceFetcher(ctx, ticker)
	
}

var priceMocks =map[string]float64{
	"BTC":12_867.0,
	"ETH": 323.0,
}

func MockPriceFetcher(ctx context.Context, ticker string)(float64, error){
	price, ok :=priceMocks[ticker]
	if!ok{
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}