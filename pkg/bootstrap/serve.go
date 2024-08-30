package bootstrap

import (
	"resedist/pkg/config"
	"resedist/pkg/html"
	"resedist/pkg/routing"
	"resedist/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
