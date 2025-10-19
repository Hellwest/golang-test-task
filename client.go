package main

import "context"

type Client struct {
	service Service
	limiter *RateLimiter
	n       uint64
	p       int64
}

func NewClient(service Service) *Client {
	n, p := service.GetLimits()

	return &Client{
		service: service,
		limiter: NewRateLimiter(n, p),
		n:       n,
		p:       p.Milliseconds(),
	}
}

func (c *Client) ProcessItems(items []Item) {
	for i := 0; i < len(items); {
		batchSize := int(c.n)

		if remaining := len(items) - i; remaining < batchSize {
			batchSize = remaining
		}

		batch := items[i : i+batchSize]

		c.limiter.Wait(uint64(len(batch)))

		c.service.Process(context.TODO(), batch)

		i += batchSize
	}
}
