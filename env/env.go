package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the configuration values
type Config struct {
	PROCESS_NAME  string
	SERVER_IP     string
	API_PORT      string
	AW_USERNAME   string
	AW_PASSWORD   string
	POS_USERNAME  string
	POS_PASSWORD  string
	DATABASE_URL  string
	CACHE_TTL     string
	LOG_LEVEL     string
	DEPL_ENV      string
	KAFKA_LOG     string
	KAFKA_BROKERS string
}

// ENV is the exported instance of Config
var ENV Config

func init() {
	// Load environment variables from .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Initialize the configuration
	ENV = Config{
		PROCESS_NAME:  getEnv("PROCESS_NAME", "default-process"),
		SERVER_IP:     getEnv("SERVER_IP", "0.0.0.0"),
		API_PORT:      getEnv("API_PORT", "8080"),
		AW_USERNAME:   getEnv("AW_USERNAME", ""),
		AW_PASSWORD:   getEnv("AW_PASSWORD", ""),
		POS_USERNAME:  getEnv("POS_USERNAME", ""),
		POS_PASSWORD:  getEnv("POS_PASSWORD", ""),
		DATABASE_URL:  getEnv("DATABASE_URL", ""),
		CACHE_TTL:     getEnv("CACHE_TTL", "15m"),
		LOG_LEVEL:     getEnv("LOG_LEVEL", "debug"),
		DEPL_ENV:      getEnv("DEPL_ENV", "DEV"),
		KAFKA_LOG:     getEnv("KAFKA_LOG", "N"),
		KAFKA_BROKERS: getEnv("KAFKA_BROKERS", "localhost:8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
