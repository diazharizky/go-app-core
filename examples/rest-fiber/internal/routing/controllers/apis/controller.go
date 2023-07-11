package apis

import (
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/app"
	"github.com/gofiber/fiber/v2"
)

type controller struct {
	appCore *app.Core
}

const traceName string = "controllers.apis"

func RegisterController(router fiber.Router, appCore *app.Core) {
	ctl := controller{appCore}
	routes := router.Group("/apis")

	routes.Get("/", ctl.List)
}
