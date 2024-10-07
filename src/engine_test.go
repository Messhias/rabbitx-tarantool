package main

import (
	"math/rand"
	"testing"
)

func TestProcessOrder(t *testing.T) {
	savePosition = func(userID int, market string, position Position) error {
		return nil
	}

	engine := NewMatchingEngine()
	order := Order{
		UserID:   rand.Int(),
		Market:   "BTC",
		Side:     "buy",
		Price:    1000,
		Size:     0.5,
		Leverage: 10,
	}

	success, err := engine.ProcessOrder(order)
	if !success || err != nil {
		t.Errorf("Expected success, but got error: %v", err)
	}

	order = Order{
		UserID:   rand.Int(),
		Market:   "BTC",
		Side:     "buy",
		Price:    1000,
		Size:     105,
		Leverage: 10,
	}

	_, err = engine.ProcessOrder(order)
	if err != nil {
		t.Errorf("Expected success, but got error: %v", err)
	}
}
