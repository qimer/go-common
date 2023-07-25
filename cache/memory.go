package cache

import (
	"context"
	"fmt"
	"time"
)

type Provider[T any] interface {
	GetData(ctx context.Context) (T, error)
}

type Cache[T any] struct {
	provider        Provider[T]
	data            T
	refreshInterval time.Duration
}

func (c *Cache[T]) Get() T {
	return c.data
}

func (c *Cache[T]) refresh() {
	if c.refreshInterval <= 0 {
		return
	}
	ctx := context.Background()
	ticker := time.NewTicker(c.refreshInterval)
	for {
		select {
		case <-ticker.C:
			t := time.Now()
			d, err := c.provider.GetData(ctx)
			if err != nil {
				fmt.Printf("refresh fail at %s, err:'%s'", t.Format(time.DateTime), err.Error())
			} else {
				fmt.Printf("refresh success at %s", t.Format(time.DateTime))
				c.data = d
			}
		}
	}
}

func NewCache[T any](ctx context.Context, refreshInterval time.Duration, provider Provider[T]) (*Cache[T], error) {
	if provider == nil {
		panic("nil provider")
	}
	data, err := provider.GetData(ctx)
	if err != nil {
		return nil, err
	}
	c := &Cache[T]{data: data, provider: provider, refreshInterval: refreshInterval}
	go c.refresh()
	return c, nil
}
