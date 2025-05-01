package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New(token)
	if err != nil {
		fmt.Println("Erro ao criar sessão do Discord:", err)
		return
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent
	dg.AddHandler(messageHandler)

	if err := dg.Open(); err != nil {
		fmt.Println("Erro ao conectar:", err)
		return
	}

	fmt.Println("Bot está online. Ctrl+C para sair.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	dg.Close()
}
