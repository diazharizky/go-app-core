package app

import "github.com/diazharizky/go-app-core/pkg/app"

type Core struct {
	// Once you are happy with the bootstrap
	// you can remove the `app.Core` code and copy the entire code inside
	// pkg/app/core.go and paste that here
	app.Core

	UserRepository IUserRepository
}
