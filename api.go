package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
)

type JSONAPIServer struct{
	svc priceFetcher
}

type PriceResponse struct{
	Ticker string `json:"ticker"`
	Price float64 `json:"price"`
}

type APIFUnc func(context.Context, http.ResponseWriter, *http.Request)error

func (s *JSONAPIServer) Run(){
	http.HandleFunc("/", )
}

func makeHTTPHandlerFunc(apiFn APIFUnc) http.HandlerFunc{
	ctx :=context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err !=nil{
			w.WriteHeader(http.StatusBadRequest)
			writeJSON(w, http.StatusBadRequest, map[string]any{"error":err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error{
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx,ticker)

	if err != nil {
		return err
	}
	priceResp := PriceResponse{
		Price:price,
		Ticker:ticker,
	}
	return writeJSON(w, http.StatusOK, &priceResp)
}

func writeJSON(w http.ResponseWriter, s int, v any) error{
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}