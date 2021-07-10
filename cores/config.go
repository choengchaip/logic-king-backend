package cores

import (
	"github.com/joho/godotenv"
	"os"
)

type IConfig interface {
	Username() string
	Password() string
	DBName() string
	DBHost() string
	DBPort() string
}

type Config struct {
	username string
	password string
	dbName   string
	dbHost   string
	dbPort   string
}

func NewConfig() IConfig {
	_ = godotenv.Load(".env")
	return &Config{}
}

func (c *Config) Username() string {
	if c.username == "" {
		c.username = os.Getenv("DB_USER")
	}

	return c.username
}

func (c *Config) Password() string {
	if c.password == "" {
		c.password = os.Getenv("DB_PASS")
	}

	return c.password
}

func (c *Config) DBName() string {
	if c.dbName == "" {
		c.dbName = os.Getenv("DB_NAME")
	}

	return c.dbName
}

func (c *Config) DBHost() string {
	if c.dbHost == "" {
		c.dbHost = os.Getenv("DB_HOST")
	}

	return c.dbHost
}

func (c *Config) DBPort() string {
	if c.dbPort == "" {
		c.dbPort = os.Getenv("DB_PORT")
	}

	return c.dbPort
}
