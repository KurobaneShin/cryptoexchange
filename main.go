package main

import (
	"fmt"

	"github.com/KurobaneShin/crypto-exchange/orderbook"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Start(":3000")

	fmt.Println("working")
}

type Market string

const(
	MarketETH Market = "ETH"
)

type Exchange struct {
	orderbooks map[Market]*orderbook.Orderbook
}

func NewExchange()*Exchange{
	return &Exchange{
		orderbooks:make(map[Market]*orderbook.Orderbook),
	}
}

func handlePlaceOrder(c echo.Context) error {
	return c.JSON(200,"SSS")
}
