package main

import (
	"fmt"
	"log"

	"github.com/diazharizky/go-app-core/config"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/app"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/repositories"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/routing"
	"github.com/diazharizky/go-app-core/pkg/rdb"
	"github.com/diazharizky/go-app-core/pkg/redix"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
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

	appCore.TracerProvider, err = tracerProvider()
	handleErr(err)

	otel.SetTracerProvider(appCore.TracerProvider)

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

func tracerProvider() (*tracesdk.TracerProvider, error) {
	url := config.Global.GetString("jaeger.url")
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, fmt.Errorf("error unable to init TracerProvider: %v", err)
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			appCore.Attributes()...,
		)),
	)

	return tp, nil
}
