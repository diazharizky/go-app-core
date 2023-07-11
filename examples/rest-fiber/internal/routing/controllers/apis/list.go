package apis

import (
	"context"
	"net/http"

	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/models"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/pkg/apiresp"
	"github.com/diazharizky/go-app-core/pkg/httpreq"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

type listRespBody struct {
	Users []models.User `json:"data"`
}

func (ctl controller) List(ctx *fiber.Ctx) error {
	_, span := otel.
		Tracer(traceName).
		Start(context.Background(), "list")

	// Set span's attributes here
	// span.SetAttributes(attribute.Key("attributes_name").String("attributes_value"))

	defer span.End()

	reqClient := httpreq.New(httpreq.HTTPReqConfig{
		BaseURL: "http://localhost:1180",
	})

	var users listRespBody
	if err := reqClient.Get("users", nil, &users); err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.FatalError())
	}

	return ctx.JSON(users)
}
