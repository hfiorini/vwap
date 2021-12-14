package ws

const (
	SubscriptionRequestType = "subscribe"
)

type Channel struct {
	Name       string
	ProductIDs []string
}

type FeedRequest struct {
	Type       string    `json:"type"`
	ProductIDs []string  `json:"product_ids"`
	Channels   []Channel `json:"channels"`
}

type FeedResponse struct {
	Type      string    `json:"type"`
	Channels  []Channel `json:"channels"`
	Message   string    `json:"message,omitempty"`
	Size      string    `json:"size"`
	Price     string    `json:"price"`
	ProductID string    `json:"product_id"`
}
