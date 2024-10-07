package main

import (
	"log"
)

// Função para salvar a posição no banco de dados Tarantool
func SavePosition(userID int, market string, position Position) error {
	_, err := tntConnection.Replace("positions", []interface{}{userID, market, position.Size, position.EntryPrice, position.Side})
	if err != nil {
		log.Printf("Error in save position: %v", err)
		return err
	}
	return nil
}
