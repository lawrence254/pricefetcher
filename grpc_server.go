package main

import (
	"context"

	"github.com/lawrence254/pricefetcher/proto"
)



type GRPCPriceFetcherServer struct{
	svc PriceFetcher
}

func NewGRPCPriceFetcherServer(svc PriceFetcher) *GRPCPriceFetcherServer{
	return &GRPCPriceFetcherServer{
		svc:svc,
	}
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest)(*proto.PriceResponse, error){
	price, err:= s.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}
	resp := &proto.PriceResponse{
		Ticker: req.Ticker,
		Price: float32(price),
	}

	return resp, err
}