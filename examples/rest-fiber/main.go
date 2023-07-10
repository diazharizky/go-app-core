package main

import (
	"log"

	"github.com/diazharizky/go-app-core/config"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/app"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/repositories"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/routing"
	pkgapp "github.com/diazharizky/go-app-core/pkg/app"
	"github.com/diazharizky/go-app-core/pkg/rdb"
	"github.com/diazharizky/go-app-core/pkg/redix"
)

var appCore *app.Core

func init() {
	config.Global.SetDefault("app.name", "go-core-app")
	config.Global.SetDefault("app.env", "development")
}

func main() {
	initAppCore()
	defer appCore.Close()

	router := routing.NewRouter(appCore)

	router.Start()
}

func initAppCore() {
	var err error

	appCore = &app.Core{}
	appCore.Info.Name = config.Global.GetString("app.name")
	appCore.Info.Version = config.Global.GetString("app.version")
	appCore.Info.Env = config.Global.GetString("app.env")

	// To set custom attributes
	// appCore.AddAttribute(attribute.String("customAttributeKey", "customAttributeValue"))

	appCore.SetTracerProvider(pkgapp.TraceExporterJaeger)

	db, err := rdb.GetDB(rdb.DBTypePostgres)
	handleErr(err)

	appCore.RDB = db

	redix, err := redix.New()
	handleErr(err)

	appCore.Redix = redix

	appCore.UserRepository = repositories.NewUserRepository(appCore.RDB)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalf("Error has happened: %v\n", err)
	}
}
