package httpreq

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	urlutils "net/url"
)

type httpReq struct {
	baseURL string
	client  *http.Client
	headers map[string]string
}

type HTTPReqConfig struct {
	BaseURL           string
	BaseHeaders       map[string]string
	Timeout           *time.Duration
	Name              string
	RateLimit         int
	RateLimitCooldown time.Duration
}

func New(cfg HTTPReqConfig) *httpReq {
	if cfg.Timeout == nil {
		defaultTimeout := time.Second * 5
		cfg.Timeout = &defaultTimeout
	}

	return &httpReq{
		baseURL: cfg.BaseURL,
		headers: cfg.BaseHeaders,
		client: &http.Client{
			Timeout: *cfg.Timeout,
		},
	}
}

func (h httpReq) Get(
	path string,
	params map[string]string,
	dest interface{},
) error {
	return h.sendRequest(http.MethodGet, path, params, nil, dest)
}

func (h httpReq) Post(
	path string,
	params map[string]string,
	body io.Reader,
	dest interface{},
) error {
	return h.sendRequest(http.MethodPost, path, params, body, dest)
}

func (h httpReq) sendRequest(
	method, path string,
	params map[string]string,
	body io.Reader,
	dest interface{},
) error {
	url := fmt.Sprintf("%s/%s", h.baseURL, path)

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

	for key, val := range h.headers {
		req.Header.Add(key, val)
	}

	resp, err := h.client.Do(req)
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

	return nil
}
