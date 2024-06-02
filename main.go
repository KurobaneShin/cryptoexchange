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

	bidParams := client.PlaceLimitOrderParams{
		UserID: 8888,
		Bid:    true,
		Price:  10_000,
		Size:   1000,
	}

	go func() {
		for {
			res, err := c.PlaceLimitOrder(&bidParams)
			if err != nil {
				panic(err)
			}

			fmt.Println("orderId =>", res.OrderID)

			time.Sleep(time.Second * 1)
		}
	}()

	askParams := client.PlaceLimitOrderParams{
		UserID: 8888,
		Bid:    false,
		Price:  8_000,
		Size:   1000,
	}

	go func() {
		for {
			res, err := c.PlaceLimitOrder(&askParams)
			if err != nil {
				panic(err)
			}

			fmt.Println("orderId =>", res.OrderID)
			time.Sleep(time.Second * 1)
		}
	}()

	select {}
}
