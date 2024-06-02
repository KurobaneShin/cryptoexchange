package main

import (
	"fmt"
	"time"

	"github.com/KurobaneShin/crypto-exchange/client"
	"github.com/KurobaneShin/crypto-exchange/server"
)

func main() {
	go server.StartServer()

	time.Sleep(time.Second * 1)

	c := client.NewClient()

	go func() {
		for {

			limitOrderParams := &client.PlaceOrderParams{
				UserID: 8888,
				Bid:    false,
				Price:  10_000,
				Size:   500_000,
			}

			res, err := c.PlaceLimitOrder(limitOrderParams)
			if err != nil {
				panic(err)
			}

			fmt.Println("placed limit order, orderId =>", res.OrderID)

			otherLimitOrderParams := &client.PlaceOrderParams{
				UserID: 8888,
				Bid:    false,
				Price:  9_000,
				Size:   500_000,
			}

			_, err = c.PlaceLimitOrder(otherLimitOrderParams)
			if err != nil {
				panic(err)
			}

			// time.Sleep(time.Second * 1)
			// if err := c.CancelOrder(res.OrderID); err != nil {
			// 	panic(err)
			// }

			marketOrderParams := &client.PlaceOrderParams{
				UserID: 7777,
				Bid:    true,
				Size:   1_000_000,
			}

			res, err = c.PlaceMarketOrder(marketOrderParams)
			if err != nil {
				panic(err)
			}

			fmt.Println("placed market order, orderId =>", res.OrderID)
			time.Sleep(time.Second * 1)
		}
	}()

	select {}
}
