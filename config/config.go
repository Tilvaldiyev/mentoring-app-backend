package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// Configuration structure
type Configuration struct {
	Server *Server `json:"server"`
	DB     *DB     `json:"db"`
}

// DB config structure
type DB struct {
	Username        string        `json:"username"`
	Name            string        `json:"name"`
	Host            string        `json:"host"`
	Port            int           `json:"port"`
	SSLMode         string        `json:"ssl_mode"`
	TimeZone        string        `json:"time_zone"`
	MaxOpenConns    int           `json:"max_open_connects"`
	MaxIdleConns    int           `json:"max_idle_connects"`
	ConnMaxLifetime time.Duration `json:"connect_max_lifetime"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time"`
	Migrate         bool          `json:"migrate"`
	LogMode         bool          `json:"log_mode"`
}

// Server config structure
type Server struct {
	Port        int  `json:"port"`
	ReleaseMode bool `json:"release_mode"`
}

// InitConfig - initializing config
func InitConfig(path string) (*Configuration, error) {
	c := new(Configuration)
	err := c.readFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("error occured while initializing Configuration: %w", err)
	}
	return c, nil
}

// LoadEnv - load environment variables from .env file
func LoadEnv() error {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return fmt.Errorf(".env file not found: %w", err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error occured while loading env variables: %w", err)
	}

	return nil
}

// readFromFile - reading data from config file and parsing it to Configuration structure
func (c *Configuration) readFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("openning file err: %w", err)
	}
	defer func() {
		err = file.Close()
	}()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		return fmt.Errorf("decoding err: %w", err)
	}

	return err
}
