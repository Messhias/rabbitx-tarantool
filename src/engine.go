package main

import (
	"errors"
	"fmt"
	"sync"
)

type MatchingEngine struct {
	Positions map[int]map[string]Position
	mu        sync.Mutex
}

func NewMatchingEngine() *MatchingEngine {
	return &MatchingEngine{
		Positions: make(map[int]map[string]Position),
	}
}

func (engine *MatchingEngine) ProcessOrder(order Order) (bool, error) {
	engine.mu.Lock()
	defer engine.mu.Unlock()

	currentPrice := 46000.0

	if engine.Positions[order.UserID] == nil {
		engine.Positions[order.UserID] = make(map[string]Position)
	}

	userPositions := engine.Positions[order.UserID]
	position, exists := userPositions[order.Market]
	if !exists {
		position = Position{Market: order.Market, Size: 0, EntryPrice: order.Price, Side: order.Side}
	}

	position.Size += order.Size
	position.EntryPrice = (position.EntryPrice + order.Price) / 2

	pnl := position.CalculatePNL(currentPrice)
	accountEquity := 100.0 + pnl
	totalNotional := position.Size * currentPrice
	margin := accountEquity / totalNotional

	if margin < 0.10 {
		return false, errors.New("No margin to do the transaction")
	}

	engine.Positions[order.UserID][order.Market] = position

	if err := savePosition(order.UserID, order.Market, position); err != nil {
		return false, err
	}

	fmt.Println("Order executed successfully: ", position)
	return true, nil
}

func (engine *MatchingEngine) StartProcessingOrders(orderChan chan Order) {
	for order := range orderChan {
		go func(ord Order) {
			_, err := engine.ProcessOrder(ord)
			if err != nil {
				fmt.Printf("Error in order processing: %v\n", err)
			}
		}(order)
	}
}
