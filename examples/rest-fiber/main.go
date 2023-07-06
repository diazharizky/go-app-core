package main

import (
	"log"

	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/app"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/repositories"
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/routing"
	"github.com/diazharizky/go-app-core/pkg/rdb"
	"github.com/diazharizky/go-app-core/pkg/redix"
)

var appCore *app.Core

func main() {
	initAppCore()
	defer appCore.Close()

	router := routing.NewRouter(appCore)

	router.Start()
}

func initAppCore() {
	appCore = &app.Core{}

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
