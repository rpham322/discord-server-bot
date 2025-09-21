package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
    "strings"
    "net/http"



	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)





func main() {
	// Load environment variables from .env
	_ = godotenv.Overload()

	token := strings.TrimSpace(os.Getenv("DISCORD_BOT_TOKEN"))
    
    
    if token == "" { log.Fatal("Token missing (.env not loaded)") }


	if token == "" {
		log.Fatal("Missing DISCORD_BOT_TOKEN in .env")
	}

	// Create Discord session
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Basic handler just to confirm the bot replies
	s.AddHandler(func(ses *discordgo.Session, ic *discordgo.InteractionCreate) {
		if ic.Type != discordgo.InteractionApplicationCommand {
			return
		}
		if ic.ApplicationCommandData().Name == "ping" {
			ses.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong! 🏓",
				},
			})
		}
	})

	// Open connection
	if err := s.Open(); err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	defer s.Close()

	log.Println("Bot is running. Try /ping in your server.")

	// Register a single test command
	cmd := &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Simple test command",
	}
	_, err = s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd)
	if err != nil {
		log.Printf("Error registering command: %v", err)
	}

	// Wait for CTRL+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down gracefully.")
}
