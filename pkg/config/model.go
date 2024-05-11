package config

var Config AppConfig

type AppConfig struct {
	App App
}

type App struct {
	Env  string
	Port string
	Name string
}

type TypeEnvironment string

const (
	LOCAL       TypeEnvironment = "local"
	DEVELOPMENT TypeEnvironment = "development"
	STAGING     TypeEnvironment = "staging"
	PRODUCTION  TypeEnvironment = "production"
)
