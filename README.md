# 🤖 Uminha Bot – Assistente para Comunidade POE2

Um bot do Discord criado para auxiliar jogadores de **Path of Exile 2**, trazendo informações atualizadas sobre preços de itens, promoções e até **dicas com IA sobre builds, farm e progressão**.

> Desenvolvido em Go com foco em performance, usando IA da GROQ e dados da API do POE.Ninja.

---

## ✨ Funcionalidades

- 📈 Consulta de **preços de itens** em tempo real.
- 💰 Lista de **promoções do dia** com melhores moedas e valores.
- 🤖 Dicas inteligentes com **IA baseada no modelo LLaMA 3.3-70B Versatile** da **GROQ**.
- 📚 Ajuda interativa com os comandos disponíveis.

---

## 📦 Tecnologias Utilizadas

- **[Go 1.22.2](https://golang.org)**
- **[discordgo](https://github.com/bwmarrin/discordgo)** – Integração com a API do Discord
- **[resty](https://github.com/go-resty/resty)** – Requisições HTTP simples e elegantes
- **[godotenv](https://github.com/joho/godotenv)** – Carregamento de variáveis de ambiente
- **[GROQ API](https://console.groq.com/)** – Consultas de IA com LLaMA 3.3
- **[POE.Ninja API](https://poe.ninja/api/data/currencyoverview?league=Necropolis&type=Currency)** – Dados atualizados de economia do jogo

---

## 💬 Comandos Disponíveis

| Comando             | Descrição                                                                 |
|---------------------|---------------------------------------------------------------------------|
| `!preco <item>`     | Mostra o preço de um item. Ex: `!preco exalted orb`                      |
| `!promo`            | Mostra as promoções do dia com os itens mais valiosos                    |
| `!dica <mensagem>`  | Recebe uma resposta de IA sobre builds, farm, economia ou progressão     |
| `!help`             | Exibe todos os comandos disponíveis                                       |

---

## 🚀 Como Rodar Localmente

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/uminha-bot.git
   cd uminha-bot
   ```

2. Crie um arquivo .env com suas variáveis de ambiente:

    ```bash
    DISCORD_TOKEN=seu_token_discord
    GROQ_API_KEY=sua_chave_groq
    ```

3. Execute o projeto:

    ```bash
    go run
    ```

## 🧠 IA Utilizada
Este bot utiliza o modelo LLaMA 3.3-70B Versatile da GROQ, que fornece respostas contextuais sobre o universo de POE2 com alto desempenho e baixa latência.


## 📌 API Externa
Para obter os preços dos itens, usamos a API oficial do POE.Ninja:

```bash
    GET https://poe.ninja/api/data/currencyoverview?league=Necropolis&type=Currency
```
