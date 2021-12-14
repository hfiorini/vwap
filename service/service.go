package service

import (
	"context"
	"strconv"
	ce "zerovwap/calculation_engine"
	"zerovwap/configuration"
	cErrors "zerovwap/custom_errors"
	"zerovwap/ws"
)

type Service interface {
	Execute(ctx context.Context) error
}

type service struct {
	wsClient ws.Client
	cfg      *configuration.Config
	items    ce.Subscription
}

func NewService(ws ws.Client, config *configuration.Config) *service {
	return &service{
		wsClient: ws,
		cfg:      config,
		items:    ce.NewSubscription(config.MaxSize),
	}
}

func (s *service) Execute(ctx context.Context) error {
	wsFeedChannel := make(chan ws.FeedResponse)
	err := s.wsClient.Subscribe(ctx, s.cfg.TradingPairs, wsFeedChannel)
	if err != nil {
		return cErrors.New("error on subscribe process", err)
	}

	for data := range wsFeedChannel {

		if data.Price == "" {
			continue
		}

		price, err := strconv.ParseFloat(data.Price, 64)
		if err != nil {
			return cErrors.New("error parsing price value to float", err)
		}
		volume, err := strconv.ParseFloat(data.Size, 64)
		if err != nil {
			return cErrors.New("error parsing volume value to float", err)
		}

		s.items.AddItem(ce.ItemSubscription{
			Price:     price,
			Volume:    volume,
			ProductID: data.ProductID,
		})

		s.items.PrintVWAP()
	}

	return nil
}
