package calculation_engine

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVWAP(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name     string
		Items    []ItemSubscription
		Expected map[string]float64
		MaxSize  int
	}{
		{
			Name:     "No items",
			Items:    []ItemSubscription{},
			Expected: map[string]float64{},
		},
		{
			Name: "Exceeding Max Size",
			Items: []ItemSubscription{
				{Price: float64(100), Volume: float64(10), ProductID: "BTC-USD"},
				{Price: float64(100), Volume: float64(20), ProductID: "ETH-USD"},
				{Price: float64(100), Volume: float64(50), ProductID: "ETH-BTC"},
				{Price: float64(100), Volume: float64(2), ProductID: "BTC-USD"},
			},
			MaxSize: 3,
			Expected: map[string]float64{
				"BTC-USD": 150,
				"ETH-USD": 100,
				"ETH-BTC": 100,
			},
		},
		{
			Name: "Happy Path",
			Items: []ItemSubscription{
				{Price: float64(100), Volume: float64(10), ProductID: "BTC-USD"},
				{Price: float64(100), Volume: float64(20), ProductID: "ETH-USD"},
				{Price: float64(100), Volume: float64(50), ProductID: "ETH-BTC"},
			},
			MaxSize: 10,
			Expected: map[string]float64{
				"BTC-USD": 100,
				"ETH-USD": 100,
				"ETH-BTC": 100,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			subscription := NewSubscription(tt.MaxSize)

			for _, item := range tt.Items {
				subscription.AddItem(item)
			}
			for k := range tt.Expected {
				require.Equal(t, tt.Expected[k], subscription.VWAPMap[k])
			}
		})
	}
}
