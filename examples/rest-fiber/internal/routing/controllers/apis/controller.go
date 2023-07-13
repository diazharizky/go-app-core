package apis

import (
	"log"
	"time"

	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/app"
	"github.com/diazharizky/go-app-core/pkg/httpreq"
	"github.com/gofiber/fiber/v2"
)

type controller struct {
	httpReq *httpreq.HTTPReq
}

const traceName string = "controllers.apis"

func RegisterController(router fiber.Router, appCore *app.Core) {
	httpr, err := httpreq.New(httpreq.HTTPReqConfig{
		BaseURL:     "http://localhost:1180",
		ServiceName: "mockserver",
		RateLimitConfig: httpreq.RateLimitConfig{
			RateLimit:         10,
			RateLimitCooldown: 60 * time.Second,
			CacheURL:          "localhost:6379",
		},
	})

	if err != nil {
		log.Fatalf("Error unable to initialize HTTP request client: %v\n", err)
	}

	ctl := controller{
		httpReq: httpr,
	}
	routes := router.Group("/apis")

	routes.Get("/", ctl.List)
}
