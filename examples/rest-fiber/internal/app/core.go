package app

import "github.com/diazharizky/go-app-core/pkg/app"

type Core struct {
	app.Core

	UserRepository IUserRepository
}
