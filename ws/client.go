package ws

import (
	"context"
	"encoding/json"
	"log"
	"net/url"
	cErrors "zerovwap/custom_errors"

	"github.com/gorilla/websocket"
	"zerovwap/configuration"
)

type Client interface {
	Subscribe(ctx context.Context, tradingPairs []string, wsFeedChannel chan FeedResponse) error
}

type client struct {
	conn *websocket.Conn
}

func NewClient(config *configuration.Config) (Client, error) {
	if config.CoinbaseUrl == "" {
		return nil, cErrors.NewWithType("WS url shouldn't be empty", cErrors.InvalidInput, nil)
	}
	u := url.URL{Scheme: "wss", Host: config.CoinbaseUrl}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, cErrors.NewWithType("Could not get a conn with WS", cErrors.Panic, err)
	}

	log.Print("Connection Success")

	return &client{
		conn: conn,
	}, nil
}

func (c *client) Subscribe(ctx context.Context, tradingPairs []string, wsFeedChannel chan FeedResponse) error {

	subscription := FeedRequest{
		Type:       SubscriptionRequestType,
		ProductIDs: tradingPairs,
		Channels: []Channel{
			{Name: "matches"},
		},
	}

	body, err := json.Marshal(subscription)
	if err != nil {
		return cErrors.New("Error marshalling subscription", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, body)
	if err != nil {
		return cErrors.New("Error sending message to WS", err)
	}

	var subscriptionResponse FeedResponse

	_, p, err := c.conn.ReadMessage()

	if err != nil {
		return cErrors.New("Error reading message from ws", err)
	}

	err = json.Unmarshal(p, &subscriptionResponse)
	if err != nil {
		return cErrors.New("Error unmarshalling subscription response", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				err := c.conn.Close()
				if err != nil {
					log.Printf("Unable to close WS Conection: %s", err)
				}
			default:
				var message FeedResponse

				_, p, err := c.conn.ReadMessage()
				if err != nil {
					log.Printf("failed receiving message: %s", err)

					break
				}

				err = json.Unmarshal(p, &message)
				if err != nil {
					log.Printf("failed unmarshalling message: %s", err)

					break
				}

				wsFeedChannel <- message

			}
		}
	}()
	return nil
}
