package bootstrap

import (
	"resedist/pkg/config"
	"resedist/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}
