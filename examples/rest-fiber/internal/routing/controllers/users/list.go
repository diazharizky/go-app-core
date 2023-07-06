package users

import (
	"net/http"

	"github.com/diazharizky/go-app-core/examples/rest-fiber/pkg/apiresp"
	"github.com/gofiber/fiber/v2"
)

func (ctl controller) List(ctx *fiber.Ctx) error {
	users, err := ctl.appCore.UserRepository.List()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(apiresp.FatalError())
	}

	return ctx.JSON(users)
}
