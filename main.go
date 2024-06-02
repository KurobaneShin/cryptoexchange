package main

import (
	"time"

	"github.com/KurobaneShin/crypto-exchange/client"
	"github.com/KurobaneShin/crypto-exchange/server"
)

func main(){
	go server.StartServer()

	time.Sleep(time.Second * 1)

	client := client.NewClient()

	if err := client.PlaceLimitOrder(); err != nil {
		panic(err)
	}
}
