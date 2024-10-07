package main

import "time"

// Estrutura de uma Ordem
type Order struct {
	UserID   int
	Market   string
	Side     string  // "buy" ou "sell"
	Price    float64 // Preço da ordem
	Size     float64 // Tamanho da posição
	Leverage int     // Alavancagem
	Time     time.Time
}

// Posição de Mercado de um Usuário
type Position struct {
	Market     string
	Size       float64
	EntryPrice float64
	Side       string
}

// Função que calcula o PNL não realizado
func (p *Position) CalculatePNL(currentPrice float64) float64 {
	sideFactor := 1.0
	if p.Side == "sell" {
		sideFactor = -1.0
	}
	return sideFactor * p.Size * (currentPrice - p.EntryPrice)
}
