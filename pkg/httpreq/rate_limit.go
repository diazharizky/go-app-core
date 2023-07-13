package httpreq

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/diazharizky/go-app-core/pkg/redix"
)

type RateLimitConfig struct {
	RateLimit         int
	RateLimitCooldown time.Duration
	CacheURL          string
}

func (h *HTTPReq) initRateLimit(cfg HTTPReqConfig) error {
	if cfg.ServiceName == "" {
		return errors.New("`ServiceName` must be configured when `RateLimit` is configured")
	}

	if cfg.RateLimitCooldown == 0 {
		return errors.New("`RateLimitCooldown` must be configured when `RateLimit` is configured")
	}

	if cfg.CacheURL == "" {
		return errors.New("`CacheURL` must be configured when `RateLimit` is configured")
	}

	var err error
	h.cache, err = redix.New(cfg.CacheURL)
	if err != nil {
		return err
	}

	return nil
}

func (h HTTPReq) checkRateLimit() error {
	if h.rateLimit <= 0 {
		return nil
	}

	key := h.rateLimitKey()
	var limit int16 = 0
	if err := h.cache.Get(context.TODO(), key, &limit); err != nil {
		return err
	}

	if limit >= h.rateLimit {
		return fmt.Errorf("%s has reached rate limit", key)
	}

	return nil
}

func (h HTTPReq) countRateLimit() error {
	if h.rateLimit <= 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var err error

	key := h.rateLimitKey()
	var limit int16 = 0
	if err = h.cache.Get(ctx, key, limit); err != nil {
		return err
	}

	limit++

	if err = h.cache.Set(ctx, key, limit); err != nil {
		return err
	}

	return nil
}

func (h HTTPReq) rateLimitKey() string {
	return "rate_limit_" + h.sanitizeServiceName()
}

func (h HTTPReq) sanitizeServiceName() string {
	sanitizedName := strings.ReplaceAll(h.serviceName, " ", "")
	return strings.ToLower(sanitizedName)
}
