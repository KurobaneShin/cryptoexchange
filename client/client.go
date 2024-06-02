package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/KurobaneShin/crypto-exchange/server"
)

const Endpoint = "http://localhost:3000"

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

func (c *Client) CancelOrder() error {
	return nil
}

type PlaceLimitOrderParams struct {
	UserID int64
	Bid    bool
	Price  float64
	Size   float64
}

func (c *Client) PlaceLimitOrder(p *PlaceLimitOrderParams) (*server.PlaceLimitOrderResponse, error) {
	params := &server.PlaceOrderRequest{
		UserID: p.UserID,
		Type:   server.LimitOrder,
		Bid:    p.Bid,
		Size:   p.Size,
		Price:  p.Price,
		Market: server.MarketETH,
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	e := Endpoint + "/order"
	req, err := http.NewRequest(http.MethodPost, e, bytes.NewReader(body))

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	placeResponse := &server.PlaceLimitOrderResponse{}

	if err := json.NewDecoder(res.Body).Decode(placeResponse); err != nil {
		return nil, err
	}

	return placeResponse, nil
}
