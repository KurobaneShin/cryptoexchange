package orderbook

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderA := NewOrder(true, 5, 0)
	buyOrderB := NewOrder(true, 8, 0)
	buyOrderC := NewOrder(true, 10, 0)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderB)

	fmt.Println(l)
}

func TestPlaceLimitOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrder1 := NewOrder(false, 10, 0)
	sellOrder2 := NewOrder(false, 5, 0)

	ob.PlaceLimitOrder(10_000, sellOrder1)
	ob.PlaceLimitOrder(9_000, sellOrder2)

	assert.Equal(t, 2, len(ob.Orders))
	assert.Equal(t, sellOrder1, ob.Orders[sellOrder1.ID])

	assert.Equal(t, len(ob.asks), 2)
}

func TestPlaceMarketOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrder := NewOrder(false, 20, 0)
	ob.PlaceLimitOrder(10_000, sellOrder)

	buyOrder := NewOrder(true, 10, 0)

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

func TestPlarceMarketOrderMultiFill(t *testing.T) {
	ob := NewOrderbook()

	buyOrderA := NewOrder(true, 5, 0)
	buyOrderB := NewOrder(true, 8, 0)
	buyOrderC := NewOrder(true, 10, 0)
	buyOrderD := NewOrder(true, 1, 0)

	ob.PlaceLimitOrder(10_000, buyOrderA)
	ob.PlaceLimitOrder(9_000, buyOrderB)
	ob.PlaceLimitOrder(5_000, buyOrderC)
	ob.PlaceLimitOrder(5_000, buyOrderD)

	assert.Equal(t, 24.00, ob.BidTotalVolume())

	sellOrder := NewOrder(false, 20, 0)
	matches := ob.PlaceMarketOrder(sellOrder)

	assert.Equal(t, ob.BidTotalVolume(), 4.0)
	assert.Equal(t, 3, len(matches))
	assert.Equal(t, 1, len(ob.bids))
}

func TestCancelOrder(t *testing.T) {
	ob := NewOrderbook()
	buyOrder := NewOrder(true, 5, 0)
	ob.PlaceLimitOrder(10_000, buyOrder)

	assert.Equal(t, 5.0, ob.BidTotalVolume())

	ob.CancelOrder(buyOrder)
	assert.Equal(t, 0.0, ob.BidTotalVolume())

	_, ok := ob.Orders[buyOrder.ID]
	assert.False(t, ok)
}
