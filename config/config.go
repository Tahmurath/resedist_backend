package config

import "time"

type Config struct {
	App     App
	Server  Server
	DB      DB
	Jsonkey Jsonkey
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

type Jsonkey struct {
	Status        string
	Error_message string
	Error_code    string
	Pagination    string
	Data          string
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
