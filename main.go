package main

import (
	"context"
	"log"
	"time"
)

var n uint64 = 20
var p = time.Second

type API struct {
	currentRequests int
}

var api = new(API)

func (a API) GetLimits() (uint64, time.Duration) {
	return n, p
}

func (a API) Process(ctx context.Context, batch Batch) error {
	time.Sleep(time.Duration(1) * time.Second)

	return nil
}

func (a *API) sendRequest(item Item) error {
	err := a.Process(context.TODO(), Batch{})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	var i uint64

	for i = 1; i <= n+5; i++ {
		log.Println("Iteration", i)
		item := Item{}

		go api.sendRequest(item)
	}
}
