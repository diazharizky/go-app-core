package users

import (
	"net/http"

	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/models"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl controller) Create(ctx *fiber.Ctx) error {
	newUser := &models.User{}

	if err := ctx.BodyParser(newUser); err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.BadRequestError())
	}

	if err := ctl.appCore.UserRepository.Create(newUser); err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.FatalError())
	}

	return ctx.JSON(newUser)

}
