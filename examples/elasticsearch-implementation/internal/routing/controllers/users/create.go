package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/diazharizky/go-app-core/examples/elasticsearch-implementation/internal/models"
	"github.com/diazharizky/go-app-core/pkg/apiresp"
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

	var newUser models.User
	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.BadRequestError())
	}

	err := ctl.appCore.UserRepository.Create(traceCtx, newUser)
	if err != nil {
		fmt.Printf("unable to create user %v:", err)

		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.FatalError())
	}

	resp := apiresp.Default()
	resp.Data = "User created"

	return ctx.JSON(resp)
}
