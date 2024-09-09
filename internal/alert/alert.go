package alert

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func Alert(message string) {
	// Load the bot token from environment variables or hardcode it (recommended to use environment variables)
	botToken := os.Getenv("DISCORD_BOT_TOKEN") // You can also hardcode your token for testing.
	if botToken == "" {
		log.Fatal("Bot token is not set in environment variables")
	}

	// Create a new Discord session using the bot token.
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Open a connection to Discord.
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection to Discord: %v", err)
	}
	defer dg.Close()

	// Replace with your channel ID
	channelID := os.Getenv("CHANNEL_ID")

	// Send a message to the specified channel.
	_, err = dg.ChannelMessageSend(channelID, message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
}
