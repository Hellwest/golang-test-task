package main

import (
	"context"
	"fmt"
	"log"
	"sync"
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
	time.Sleep(time.Duration(5) * time.Second)

	return nil
}

func (a *API) sendRequest(item Item, wg *sync.WaitGroup) error {
	fmt.Println("Here")
	err := a.Process(context.TODO(), Batch{})

	wg.Done()
	return err
}

func main() {
	var i uint64
	var wg sync.WaitGroup

	for i = 1; i <= n+5; i++ {
		log.Println("Iteration", i)
		item := Item{}

		wg.Add(1)
		go api.sendRequest(item, &wg)
	}

	wg.Wait()
}
