package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/chucky-1/broker/protocol"
	"github.com/chucky-1/trader/internal/config"
	"github.com/chucky-1/trader/internal/grpc/client"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"bufio"
	"context"
	"fmt"
	"os"
)

const instruction = "--- Input 1 to open position; 2 to close position; 3 to get balance; 4 to change balance"

func main() {
	// Configuration
	cfg := new(config.Config)
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("%v", err)
	}

	ctx := context.Background()

	// Grpc Broker
	hostAndPort := fmt.Sprint("localhost", ":", "11000")
	conn, err := grpc.Dial(hostAndPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
	cl := protocol.NewBrokerClient(conn)
	trd := client.NewTrader(cl, 4)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(instruction)
			input, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
			if err != nil {
				log.Error(err)
				continue
			}
			switch {
			case string(input[:len(input)-2]) == "1":
				id, err := trd.OpenPosition(ctx)
				if err != nil {
					log.Error(err)
				} else {
					fmt.Println(id)
				}
			case string(input[:len(input)-2]) == "2":
				err = trd.ClosePosition(ctx)
				if err != nil {
					log.Error(err)
				} else {
					fmt.Println("Position closed successfully")
				}
			case string(input[:len(input)-2]) == "3":
				sum, err := trd.GetBalance(ctx)
				if err != nil {
					log.Error(err)
				} else {
					fmt.Printf("Your balance is %f\n", sum)
				}
			case string(input[:len(input)-2]) == "4":
				err = trd.SetBalance(ctx)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
}