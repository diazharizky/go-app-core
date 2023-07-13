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

func (c clientRate) checkRateThreshold(rateKey string) error {
	if c.limit <= 0 {
		return nil
	}

	var currentRate int16 = 0
	if err := c.cache.Get(context.TODO(), rateKey, &currentRate); err != nil {
		return err
	}

	if currentRate >= c.limit {
		return fmt.Errorf("rate limit exceeded: %s", rateKey)
	}

	return nil
}

func (c clientRate) increaseRate(key string) error {
	if c.limit <= 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error

	var currentRate int16 = 0
	if err = c.cache.Get(ctx, key, &currentRate); err != nil {
		return err
	}

	currentRate++

	if err = c.cache.Set(ctx, key, currentRate); err != nil {
		return err
	}

	return nil
}
