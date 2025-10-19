package main

import (
	"context"
	"fmt"
	"time"
)

const itemsToSend = 117

type MockService struct {
	n uint64
	p time.Duration
}

func (m *MockService) GetLimits() (n uint64, p time.Duration) {
	return m.n, m.p
}

func (m *MockService) Process(ctx context.Context, batch Batch) error {
	fmt.Printf("Processed batch of %d items at %v\n", len(batch), time.Now().Format("15:04:05.000"))
	return nil
}

func main() {
	service := &MockService{n: 10, p: time.Millisecond * 500}
	client := NewClient(service)
	items := make([]Item, itemsToSend)

	client.ProcessItems(items)
}
