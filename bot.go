package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	content := strings.ToLower(m.Content)

	if content == "!promo" {
		response := BuscarTopPromos()
		s.ChannelMessageSend(m.ChannelID, response)
		return
	}

	if strings.HasPrefix(content, "!preco ") {
		item := strings.TrimSpace(strings.TrimPrefix(content, "!preco"))
		if item == "" {
			s.ChannelMessageSend(m.ChannelID, "Por favor, informe um item. Exemplo: `!preco exalted orb`")
			return
		}

		preco := BuscarPrecoItem(item)
		s.ChannelMessageSend(m.ChannelID, preco)
		return
	}
}
