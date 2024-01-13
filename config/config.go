package config

import (
	"os"
	"strconv"
)

// declare a constant for db config and server config
const (
	PulsarUrl         = "PULSAR_URL"
	PulsarServiceName = "PULSAR_SERVICE_NAME"
	ServerProtocol    = "SERVER_NAME"
	ServerPort        = "SERVER_PORT"
	SendGridApiKey    = "SENDGRID_API_KEY"
)

// Config holds fields for configuration
type Config struct {
	PulsarUrl         string
	PulsarServiceName string
	ServerProtocol    string
	ServerPort        string
	SendgridApiKey    string
}

func (c Config) GetEnv(key, fallback string) string {
	res, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}
	return res
}

func (c Config) GetEnvBool(key string, fallback bool) bool {
	res := c.GetEnv(key, "")
	if len(res) == 0 {
		return fallback
	}
	v, err := strconv.ParseBool(res)
	if err != nil {
		return fallback
	}
	return v
}

func (c Config) GetEnvInt(key string, fallback int) int {
	res, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}
	v, err := strconv.Atoi(res)
	if err != nil {
		return fallback
	}
	return v
}

func ImportConfig(c Config) *Config {
	pulsarUrl := c.GetEnv(PulsarUrl, "pulsar://localhost:6650")
	serverprotocol := c.GetEnv(ServerProtocol, "tcp")
	serverPort := c.GetEnv(ServerPort, "50053")
	sendgridKey := c.GetEnv(SendGridApiKey, "G.C3HC0xIiRC2FhzD1Iwl5ag.a6biEm3GcP77eqAVWlYJmB7lHbYFUx-GhLOgZkhBjlw")

	return &Config{
		PulsarUrl:      pulsarUrl,
		ServerProtocol: serverprotocol,
		ServerPort:     serverPort,
		SendgridApiKey: sendgridKey,
	}
}
