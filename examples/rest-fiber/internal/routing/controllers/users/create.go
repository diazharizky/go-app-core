package users

import (
	"context"
	"net/http"

	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/models"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

func (ctl controller) Create(ctx *fiber.Ctx) error {
	traceCtx, span := otel.
		Tracer(traceName).
		Start(context.Background(), "create")

	// Set span's attributes here
	// span.SetAttributes(attribute.Key("attributes_name").String("attributes_value"))

	defer span.End()

	newUser := &models.User{}

	if err := ctx.BodyParser(newUser); err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.BadRequestError())
	}

	if err := ctl.appCore.UserRepository.Create(traceCtx, newUser); err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.FatalError())
	}

	return ctx.JSON(newUser)

}
