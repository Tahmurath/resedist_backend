package bootstrap

import (
	"resedist/pkg/config"
	"resedist/pkg/database"
	"resedist/pkg/html"
	"resedist/pkg/routing"
	"resedist/pkg/sessions"
	"resedist/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	routing.ConfigureCorsConfig()

	sessions.Start(routing.GetRouter())

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
