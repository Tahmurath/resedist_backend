package config

import "time"

type Config struct {
	App     App
	Server  Server
	DB      DB
	URLKeys Urlkeys
	Dblog   Dblog
	Redis   Redis
	Log     Log
	Jwt     JWT
	Cors    CORS
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

type Urlkeys struct {
	Expand   string
	Sort     string
	Order    string
	Page     string
	Pagesize string
}

type Dblog struct {
	LogLevel int
	Colorful bool
}

type Log struct {
	LogLevel int
}

type Redis struct {
	Addr     string
	Password string
	DB       int
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
