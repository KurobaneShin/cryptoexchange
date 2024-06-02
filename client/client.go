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

func (c *Client) CancelOrder(orderID int64) error {
	e := fmt.Sprintf("%s/order/%d", Endpoint, orderID)
	req, err := http.NewRequest(http.MethodDelete, e, nil)
	if err != nil {
		return err
	}
	_, err = c.Do(req)
	if err != nil {
		return err
	}
	return nil
}

type PlaceOrderParams struct {
	UserID int64
	Bid    bool
	// Price only needed when placing LIMIT orders
	Price float64
	Size  float64
}

func (c *Client) PlaceMarketOrder(p *PlaceOrderParams) (*server.PlaceOrderResponse, error) {
	params := &server.PlaceOrderRequest{
		UserID: p.UserID,
		Type:   server.MarketOrder,
		Bid:    p.Bid,
		Size:   p.Size,
		Market: server.MarketETH,
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	e := Endpoint + "/order"
	req, err := http.NewRequest(http.MethodPost, e, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	placeResponse := &server.PlaceOrderResponse{}

	if err := json.NewDecoder(res.Body).Decode(placeResponse); err != nil {
		return nil, err
	}

	return placeResponse, nil
}

func (c *Client) PlaceLimitOrder(p *PlaceOrderParams) (*server.PlaceOrderResponse, error) {
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

	placeResponse := &server.PlaceOrderResponse{}

	if err := json.NewDecoder(res.Body).Decode(placeResponse); err != nil {
		return nil, err
	}

	return placeResponse, nil
}
