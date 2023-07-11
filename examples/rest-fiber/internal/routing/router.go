package routing

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/diazharizky/go-app-core/config"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/app"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/routing/controllers/apis"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/routing/controllers/users"
	"github.com/gofiber/fiber/v2"
)

type router struct {
	server *fiber.App
}

func init() {
	config.Global.SetDefault("rest.fiber.server.port", 18080)
}

func NewRouter(appCore *app.Core) (r router) {
	svr := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	api := svr.Group("api")
	v1 := api.Group("v1")

	users.RegisterController(v1, appCore)
	apis.RegisterController(v1, appCore)

	r.server = svr

	return
}

func (r router) Start() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nShutting down gracefully...")
		r.server.Shutdown()
	}()

	addr := fmt.Sprintf(":%d", config.Global.GetInt("rest.fiber.server.port"))
	if err := r.server.Listen(addr); err != nil {
		log.Panic(err)
	}
}
