package users

import (
	"context"
	"net/http"

	"github.com/diazharizky/go-app-core/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

func (ctl controller) List(ctx *fiber.Ctx) error {
	traceCtx, span := otel.
		Tracer(traceName).
		Start(context.Background(), "list")

	// Set span's attributes here
	// span.SetAttributes(attribute.Key("attributes_name").String("attributes_value"))

	defer span.End()

	users, err := ctl.appCore.UserRepository.List(traceCtx)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.FatalError())
	}

	resp := apiresp.Default()
	resp.Data = users

	return ctx.JSON(resp)
}
