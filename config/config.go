package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type Config struct {
	DbConfig
}

func (c *Config) ReadConfig() error {

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error load .env file")
	}

	// Environment Variable
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	// Cek jika tidak mengirimkan
	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Driver == "" {
		return fmt.Errorf("missing ENV")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

