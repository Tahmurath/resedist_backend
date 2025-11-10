package config

import "time"

type Config struct {
	App      App
	Server   Server
	DB       DB
	Rest     Rest
	Telegram Telegram
	Dblog    Dblog
	Redis    Redis
	Log      Log
	Jwt      JWT
	Cors     CORS
}

type App struct {
	Name string
}

type Server struct {
	Host           string
	Port           string
	Ginmode        string
	TrustedProxies []string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type Rest struct {
	Status        string
	Error_message string
	Error_code    string
	Pagination    string
	Data          string
	Request       string
	Success       string
	Failed        string
	Bind_error    string
	Not_found     string
}

type Telegram struct {
	BotToken  string
	TokenExpr time.Duration
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
	Secret          string
	Duration        time.Duration
	RefreshDuration time.Duration
	AccessDuration  time.Duration
}

type CORS struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
}
