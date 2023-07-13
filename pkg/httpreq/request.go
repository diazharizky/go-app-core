package httpreq

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	urlutils "net/url"

	"github.com/diazharizky/go-app-core/pkg/redix"
)

type HTTPReq struct {
	serviceName string
	baseURL     string
	headers     map[string]string
	client      *http.Client
	cache       *redix.Redix
	rateLimit   int16
}

type HTTPReqConfig struct {
	RateLimitConfig

	BaseURL     string
	BaseHeaders map[string]string
	Timeout     *time.Duration
	ServiceName string
}

func New(cfg HTTPReqConfig) (*HTTPReq, error) {
	if cfg.Timeout == nil {
		defaultTimeout := time.Second * 5
		cfg.Timeout = &defaultTimeout
	}

	httpr := &HTTPReq{
		baseURL: cfg.BaseURL,
		headers: cfg.BaseHeaders,
		client: &http.Client{
			Timeout: *cfg.Timeout,
		},
	}

	if cfg.RateLimit > 0 {
		if err := httpr.initRateLimit(cfg); err != nil {
			return nil, err
		}
	}

	return httpr, nil
}

func (h HTTPReq) Get(
	path string,
	params map[string]string,
	dest interface{},
) error {
	return h.sendRequest(http.MethodGet, path, params, nil, dest)
}

func (h HTTPReq) Post(
	path string,
	params map[string]string,
	body io.Reader,
	dest interface{},
) error {
	return h.sendRequest(http.MethodPost, path, params, body, dest)
}

func (h HTTPReq) sendRequest(
	method, path string,
	params map[string]string,
	body io.Reader,
	dest interface{},
) error {
	if err := h.checkRateLimit(); err != nil {
		return err
	}

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

	if err = h.countRateLimit(); err != nil {
		log.Printf("Error unable record rate limit count: %v\n", err)
	}

	return nil
}
