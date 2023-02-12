// Package config ...
package config

// Config for env values
type Config struct {
	PostgresDBURL   string `env:"POSTGRES_DB_URL"`
	CookieTokenName string `env:"COOKIE_TOKEN_NAME"`
	JWTKey          string `env:"JWT_KEY"`
}
