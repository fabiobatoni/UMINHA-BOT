package main

import (
	"strings"
	"uminha-bot/openai"

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

	if strings.HasPrefix(m.Content, "!dica") {
		pergunta := strings.TrimPrefix(m.Content, "!poe2 ")
		resposta, err := openai.AskGPT(pergunta)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Erro ao consultar o especialista POE2: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, resposta)
	}

	if strings.HasPrefix(m.Content, "!help") {
		helpMessage := "**📜 Comandos disponíveis:**\n" +
			"`!dica <sua pergunta>` – Pergunte algo sobre builds, classes, dicas ou economia do Path of Exile 2.\n" +
			"`!ping` – Testa se o bot está online.\n" +
			"`!help` – Mostra esta mensagem de ajuda." +
			"`!promo` – Mostra as promoções do dia.\n" +
			"`!preco <item>` – Mostra o preço de um item específico. Exemplo: `!preco exalted orb`.\n"

		s.ChannelMessageSend(m.ChannelID, helpMessage)
		return
	}

}
