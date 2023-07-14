package httpclient

import (
	"context"
	"fmt"
	"time"

	"github.com/diazharizky/go-app-core/pkg/redix"
)

type clientRate struct {
	cache    *redix.Redix
	limit    int16
	cooldown time.Duration
}

type ClientRateConfig struct {
	Limit    int16
	Cooldown time.Duration
	CacheURL string
}

type rateRecord struct {
	Counter   int16     `redis:"counter"`
	UpdatedAt time.Time `redis:"updated_at"`
}

func (c clientRate) checkRateLimit(key string) (*rateRecord, error) {
	if c.limit <= 0 {
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var currentRate rateRecord
	if err := c.cache.Get(ctx, key, &currentRate); err != nil {
		return nil, err
	}

	currentTime := time.Now()
	if currentTime.Sub(currentRate.UpdatedAt) >= c.cooldown {
		currentRate.Counter = 0
	}

	if currentRate.Counter >= c.limit {
		return nil, fmt.Errorf("rate limit exceeded: %s", key)
	}

	return &currentRate, nil
}

func (c clientRate) incrementRate(key string, rate rateRecord) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rate.Counter++
	rate.UpdatedAt = time.Now()

	if err := c.cache.Set(ctx, key, rate); err != nil {
		return err
	}

	return nil
}
