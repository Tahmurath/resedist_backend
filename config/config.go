package config

type Config struct {
	App    App
	Server Server
	DB     DB
	Dblog  Dblog
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

type JWT struct {
	Secret string
}

type CORS struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
}
