package configuration

const (
	CoinbaseURL = "ws-feed.exchange.coinbase.com"
)

type Config struct {
	CoinbaseUrl  string
	TradingPairs []string
	MaxSize      int
}

func InitConfig() *Config {
	return &Config{
		CoinbaseUrl:  CoinbaseURL,
		TradingPairs: []string{"BTC-USD", "ETH-USD", "ETH-BTC"},
		MaxSize:      200,
	}
}
