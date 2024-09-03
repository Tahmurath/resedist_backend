package bootstrap

import (
	"resedist/pkg/config"
	"resedist/pkg/database"
	"resedist/pkg/html"
	"resedist/pkg/routing"
	"resedist/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
