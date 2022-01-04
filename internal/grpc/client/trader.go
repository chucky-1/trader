// Package client has all about client of GRPC
package client

import (
	"github.com/chucky-1/broker/protocol"

	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
)

// Trader is struct
type Trader struct {
	client protocol.BrokerClient
	id int32
}

// NewTrader is constructor
func NewTrader(client protocol.BrokerClient, id int32) *Trader {
	return &Trader{client: client, id: id}
}

// OpenPosition opens a position
func (t *Trader) OpenPosition(ctx context.Context) (int32, error) {
	fmt.Print("Input id of symbol: ")
	symbolID, err := input()
	if err != nil {
		return 0, err
	}

	fmt.Print("Input the expected price: ")
	price, err := input()
	if err != nil {
		return 0, err
	}

	fmt.Print("Input amount of stock that you want to bay: ")
	count, err := input()
	if err != nil {
		return 0, err
	}

	fmt.Print("Define a stop loss if you like: ")
	stopLoss, err := input()
	if err != nil {
		return 0, err
	}

	fmt.Print("Define a take profit if you like: ")
	takeProfit, err := input()
	if err != nil {
		return 0, err
	}

	fmt.Print("Input 1 or 2. 1 for Buy, 2 for Sell: ")
	i, err := input()
	if err != nil {
		return 0, err
	}
	var isBuy bool
	if i == 1 {
		isBuy = true
	} else {
		isBuy = false
	}

	response, err := t.client.OpenPosition(ctx, &protocol.OpenPositionRequest{
		UserId: t.id,
		SymbolId: symbolID,
		Price: float32(price),
		Count: count,
		StopLoss: float32(stopLoss),
		TakeProfit: float32(takeProfit),
		IsBuy: isBuy,
	})
	if err != nil {
		return 0, err
	}
	return response.PositionId, nil
}

// ClosePosition closes a position
func (t *Trader) ClosePosition(ctx context.Context) error {
	fmt.Println("Input id of position")
	id, err := input()
	if err != nil {
		return err
	}
	_, err = t.client.ClosePosition(ctx, &protocol.ClosePositionRequest{
		PositionId: id,
	})
	if err != nil {
		return err
	}
	return nil
}

// SetBalance changes a balance
func (t *Trader) SetBalance(ctx context.Context) error {
	fmt.Println("Input sum")
	sum, err := input()
	if err != nil {
		return err
	}
	_, err = t.client.SetBalance(ctx, &protocol.SetBalanceRequest{
		UserId: t.id,
		Sum:    float32(sum),
	})
	if err != nil {
		return err
	}
	return nil
}

// GetBalance returns a balance
func (t *Trader) GetBalance(ctx context.Context) (float32, error) {
	response, err := t.client.GetBalance(ctx, &protocol.GetBalanceRequest{
		UserId: t.id,
	})
	if err != nil {
		return 0, err
	}
	return response.Sum, nil
}

func input() (int32, error) {
	e, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		return 0, err
	}
	element, err := strconv.Atoi(string(e[:len(e) - 2]))
	if err != nil {
		return 0, err
	}
	return int32(element), nil
}
