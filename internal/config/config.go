package config

import (
	"log"
	"os"
)

// Config struct holds all the configuration settings
type Config struct {
	APIURL          string
	DBConnStr       string
	DiscordBotToken string
	ChannelID       string
}

// AppConfig is the global configuration variable that will hold the loaded settings
var AppConfig *Config

// LoadConfig initializes the configuration by reading environment variables
func LoadConfig() *Config {
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		log.Fatal("API_URL environment variable is required")
	}

	dbConnStr := os.Getenv("DB_CONN_STR")
	if dbConnStr == "" {
		log.Fatal("DB_CONN_STR environment variable is required")
	}

	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")
	if discordBotToken == "" {
		log.Fatal("DISCORD_BOT_TOKEN environment variable is required")
	}

	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		log.Fatal("CHANNEL_ID environment variable is required")
	}

	AppConfig = &Config{
		APIURL:          apiURL,
		DBConnStr:       dbConnStr,
		DiscordBotToken: discordBotToken,
		ChannelID:       channelID,
	}

	return AppConfig
}
