package bootstrap

import (
	"resedist/pkg/applog"
	"resedist/pkg/config"
	"resedist/pkg/database"
	"resedist/pkg/html"
	"resedist/pkg/redis"
	"resedist/pkg/routing"
	"resedist/pkg/sessions"
	"resedist/pkg/static"
)

func Serve() {
	applog.Info("config set")
	config.Set()

	applog.Info("database connect")
	database.Connect()

	applog.Info("redis connect")
	redis.Connect()

	applog.Info("routing init")
	routing.Init()

	applog.Info("cors config")
	routing.ConfigureCorsConfig()

	applog.Info("session start")
	sessions.Start(routing.GetRouter())

	applog.Info("static load")
	static.LoadStatic(routing.GetRouter())

	applog.Info("html load")
	html.LoadHTML(routing.GetRouter())

	applog.Info("route register")
	routing.RegisterRoutes()

	applog.Info("start serve")
	routing.Serve()
}
