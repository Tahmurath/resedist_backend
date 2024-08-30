package bootstrap

import (
	"resedist/pkg/config"
	"resedist/pkg/html"
	"resedist/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
