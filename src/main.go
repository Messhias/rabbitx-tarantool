package main

import (
	"github.com/tarantool/go-tarantool"
	"log"
	"math/rand"
	"time"
)

var tntConnection *tarantool.Connection

var savePosition = func(userID int, market string, position Position) error {
	_, err := tntConnection.Replace("positions", []interface{}{userID, market, position.Size, position.EntryPrice, position.Side})
	return err
}

func main() {
	var err error
	tntConnection, err = tarantool.Connect("tarantool:3301", tarantool.Opts{
		User: "admin",
		Pass: "password",
	})
	if err != nil {
		log.Fatalf("Can't connect at tarantool: %v", err)
	}
	defer func(tntConnection *tarantool.Connection) {
		err := tntConnection.Close()
		if err != nil {
			log.Println(err)
		}
	}(tntConnection)
	log.Println("DB connected")

	numOrders := 10

	engine := NewMatchingEngine()
	orderChan := make(chan Order, numOrders)
	go engine.StartProcessingOrders(orderChan)

	for i := 0; i < numOrders; i++ {
		order := generateRandomOrder(i)
		orderChan <- order
	}

	time.Sleep(5 * time.Second)
}

func generateRandomOrder(userID int) Order {
	markets := []string{"BTC", "ETH"}
	sides := []string{"buy", "sell"}

	return Order{
		UserID:   userID,
		Market:   markets[rand.Intn(len(markets))],
		Side:     sides[rand.Intn(len(sides))],
		Price:    1000 + rand.Float64()*50000,
		Size:     rand.Float64() * 5,
		Leverage: 1 + rand.Intn(10),
		Time:     time.Now(),
	}
}
