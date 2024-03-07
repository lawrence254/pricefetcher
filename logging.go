package main

import (
	"context"
	"log/slog"
	"time"
)

type loggingService struct{
	next PriceFetcher
}

func NewLoggingService(next PriceFetcher) PriceFetcher{
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string)(price float64, err error){
	defer func(begin time.Time) {
		if err != nil {
			slog.Error("fetchPrice",
				slog.Duration("took", time.Since(begin)),
				slog.String("err", err.Error()),
			)
		} else {
			slog.Info("fetchPrice",
				slog.Int("requestID", ctx.Value("requestID").(int)),
				slog.Duration("took", time.Since(begin)),
				slog.Float64("price", price),
			)
		}
	}(time.Now())
	return s.next.FetchPrice(ctx, ticker)
}