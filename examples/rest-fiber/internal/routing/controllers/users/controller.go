package users

import (
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/app"
	"github.com/gofiber/fiber/v2"
)

type controller struct {
	appCore *app.Core
}

func RegisterController(router fiber.Router, appCore *app.Core) {
	ctl := controller{appCore}
	routes := router.Group("/users")

	routes.Get("/", ctl.List)
	routes.Post("/", ctl.Create)
}
