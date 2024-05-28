package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 8)
	buyOrderC := NewOrder(true, 10)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderB)

	fmt.Println(l)
}

func TestPlaceLimitOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrder1 := NewOrder(false, 10)
	sellOrder2 := NewOrder(false, 5)

	ob.PlaceLimitOrder(10_000, sellOrder1)
	ob.PlaceLimitOrder(9_000, sellOrder2)

	assert.Equal(t, len(ob.asks), 2)
}

func TestPlaceMarketOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrder := NewOrder(false, 20)
	ob.PlaceLimitOrder(10_000, sellOrder)

	buyOrder := NewOrder(true, 10)

	matches := ob.PlaceMarketOrder(buyOrder)

	assert.Equal(t, 1, len(matches))
	assert.Equal(t, 1, len(ob.asks))
	assert.Equal(t, 10.0, ob.AskTotalVolume())

	assert.Equal(t, sellOrder, matches[0].Ask)
	assert.Equal(t, buyOrder, matches[0].Bid)
	assert.Equal(t, 10.0, matches[0].SizeFilled)
	assert.Equal(t, 10_000.00, matches[0].Price)

	assert.True(t, buyOrder.IsFilled())
}
