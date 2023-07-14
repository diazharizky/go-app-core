package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	urlutils "net/url"

	"github.com/diazharizky/go-app-core/pkg/redix"
)

type Client struct {
	baseURL string
	headers map[string]string
	agent   *http.Client
	rate    clientRate
	apiName string
}

type ClientConfig struct {
	ClientRateConfig

	BaseURL string
	Headers map[string]string
	Timeout time.Duration
	APIName string
}

func New(cfg ClientConfig) (*Client, error) {
	if cfg.Timeout == 0 {
		cfg.Timeout = 5 * time.Second
	}

	client := &Client{
		baseURL: cfg.BaseURL,
		headers: cfg.Headers,
		agent: &http.Client{
			Timeout: cfg.Timeout,
		},
	}

	if cfg.APIName != "" {
		client.apiName = cfg.APIName
	}

	if cfg.Limit > 0 {
		if err := client.initClientRate(cfg); err != nil {
			return nil, fmt.Errorf("unable to initialize client rate config: %v", err)
		}
	}

	return client, nil
}

func (h Client) Get(
	path string,
	params map[string]string,
	dest interface{},
) error {
	return h.sendRequest(http.MethodGet, path, params, nil, dest)
}

func (h Client) Post(
	path string,
	params map[string]string,
	body io.Reader,
	dest interface{},
) error {
	return h.sendRequest(http.MethodPost, path, params, body, dest)
}

func (c Client) sendRequest(
	method, path string,
	params map[string]string,
	body io.Reader,
	dest interface{},
) error {
	currentRate, err := c.rate.checkRateLimit(c.rateKey())
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s", c.baseURL, path)

	if len(params) > 0 {
		qs := urlutils.Values{}
		for key, val := range params {
			qs.Add(key, val)
		}

		url += qs.Encode()
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	for key, val := range c.headers {
		req.Header.Add(key, val)
	}

	resp, err := c.agent.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Printf("Error unable to close response's body: %v\n", err)
		}
	}()

	if err = json.NewDecoder(resp.Body).Decode(dest); err != nil {
		return err
	}

	if currentRate != nil {
		if err = c.rate.incrementRate(c.rateKey(), *currentRate); err != nil {
			log.Printf("Error happened on increment rate count: %v\n", err)
		}
	}

	return nil
}

func (c *Client) initClientRate(cfg ClientConfig) error {
	if cfg.Cooldown == 0 {
		return errors.New("`Cooldown` must be configured when using rate limiting")
	}

	if cfg.CacheURL == "" {
		return errors.New("`CacheURL` must be configured when using rate limiting")
	}

	c.rate = clientRate{
		limit:    cfg.Limit,
		cooldown: cfg.Cooldown,
	}

	var err error
	c.rate.cache, err = redix.New(cfg.CacheURL)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) rateKey() string {
	return "api_rate_" + c.apiName
}
