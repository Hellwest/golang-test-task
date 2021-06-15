package main

import (
	"context"
	"sync"
	"time"
)

var n uint64 = 20
var p = 5 * time.Second

type Limits struct {
	n uint64
	p time.Duration
}

type API struct{}

var api = new(API)

func (a *API) GetLimits() (uint64, time.Duration) {
	return n, p
}

func (a *API) Process(ctx context.Context, batch Batch) error {
	time.Sleep(time.Duration(3) * time.Second)
	// log.Println("Returning a result")

	return nil
}

func sleepForCooldown(earliestRequestSentAt time.Time) {
	sleepTime := time.Until(earliestRequestSentAt.Add(p))

	// log.Println("Sleeping for", sleepTime)

	time.Sleep(sleepTime)
}

func sendRequests(limits Limits) {
	var wg sync.WaitGroup
	var totalBatches int = 5
	var earliestRequestSentAt time.Time

	batchNumber := 1
	for batchNumber <= totalBatches {
		if time.Since(earliestRequestSentAt) >= p {
			earliestRequestSentAt = time.Now() // TODO: Make it correlate with the actual request sending timestamp

			sendBatch(limits, &wg, batchNumber)

			batchNumber++
		} else {
			sleepForCooldown(earliestRequestSentAt)
		}
	}

	wg.Wait()
}

func sendBatch(limits Limits, wg *sync.WaitGroup, batchNumber int) {
	for i := uint64(1); i <= limits.n; i++ {
		// log.Println("Batch", batchNumber, ", sending request #", i)
		item := Item{}

		wg.Add(1)
		go sendRequest(item, wg)
	}
}

func sendRequest(item Item, wg *sync.WaitGroup) error {
	// log.Println("Sending")
	defer wg.Done()

	return api.Process(context.TODO(), Batch{})
}

func main() {
	n, p := api.GetLimits()
	limits := Limits{
		n: n,
		p: p,
	}

	sendRequests(limits)
}
