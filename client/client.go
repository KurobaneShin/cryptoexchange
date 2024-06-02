package client

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (c *Client) PlaceLimitOrder() error {
	params := &server.PlaceOrderRequest{
		UserID: 8888,
		Type:   server.LimitOrder,
		Bid:    true,
		Size:   4000.0,
		Price:  4000.0,
		Market: server.MarketETH,
	}

	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	e := Endpoint + "/order"
	req, err := http.NewRequest(http.MethodPost, e, bytes.NewReader(body))

	res,err := c.Do(req)

	if err != nil {
		return err
	}

	fmt.Printf("%+v", res)

	return nil
}
