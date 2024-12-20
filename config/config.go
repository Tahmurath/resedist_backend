package config

import "time"

type Config struct {
	App    App
	Server Server
	DB     DB
	Dblog  Dblog
	Redis  Redis
	Jwt    JWT
	Cors   CORS
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type Dblog struct {
	LogLevel int
	Colorful bool
}

type Redis struct {
	Addr     string
	Password string
	DB       string
	//Protocol     string
}

type JWT struct {
	Secret   string
	Duration time.Duration
}

type CORS struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
}
