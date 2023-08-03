package main

import (
	"log"

	"github.com/diazharizky/go-app-core/config"
	"github.com/diazharizky/go-app-core/examples/elasticsearch-implementation/internal/app"
	"github.com/diazharizky/go-app-core/examples/elasticsearch-implementation/internal/repositories"
	"github.com/diazharizky/go-app-core/examples/elasticsearch-implementation/internal/routing"
	pkgapp "github.com/diazharizky/go-app-core/pkg/app"
	"github.com/diazharizky/go-app-core/pkg/elasticsearch"
)

var appCore *app.Core

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

	esClient, err := elasticsearch.GetClient()
	handleErr(err)

	appCore.UserRepository = repositories.NewUserRepository(esClient)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalf("Error has happened: %v\n", err)
	}
}
