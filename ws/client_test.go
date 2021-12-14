package ws

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"zerovwap/configuration"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		Name      string
		Config    *configuration.Config
		WantError bool
	}{
		{
			Name: "Empty Coinbase Url should fail",
			Config: &configuration.Config{
				CoinbaseUrl: "",
			},
			WantError: true,
		},
		{
			Name: "Success",
			Config: &configuration.Config{
				CoinbaseUrl: "ws-feed.exchange.coinbase.com",
			},
			WantError: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			_, err := NewClient(tt.Config)
			if tt.WantError {
				if err == nil {
					t.Fail()
				}
			} else {
				if err != nil {
					t.Fail()
				}
			}
		})
	}
}

func TestClient_Subscribe(t *testing.T) {
	tests := []struct {
		Name      string
		Config    *configuration.Config
		WantError bool
	}{
		{
			Name: "Empty Coinbase Url should fail",
			Config: &configuration.Config{
				CoinbaseUrl:  "",
				TradingPairs: []string{"ETH-BTC"},
			},
			WantError: true,
		},
		{
			Name: "Success",
			Config: &configuration.Config{
				CoinbaseUrl:  "ws-feed.exchange.coinbase.com",
				TradingPairs: []string{"ETH-BTC"},
			},
			WantError: false,
		},
	}
	wsFeedChannel := make(chan FeedResponse)
	for _, tt := range tests {
		tt := tt

		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			ws, e := NewClient(tt.Config)
			if tt.WantError && e == nil {
				t.Fail()
			}
			if !tt.WantError && e != nil {
				t.Fail()
			}

			if ws != nil {
				err := ws.Subscribe(ctx, tt.Config.TradingPairs, wsFeedChannel)

				if tt.WantError {
					require.Error(t, err)
				} else {
					require.NoError(t, err)

					var counter int

					for m := range wsFeedChannel {
						if counter >= 2 {
							break
						}
						require.Equal(t, m.ProductID, "ETH-BTC")
						require.Equal(t, m.Message, "")

						counter++
					}
				}
			}

		})
	}
}
