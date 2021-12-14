package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitConfig(t *testing.T) {
	config := InitConfig()
	a := assert.New(t)
	a.NotEmpty(t, config.CoinbaseUrl)
	a.NotNil(t, config.TradingPairs)
	a.NotZero(t, config.MaxSize)
}
