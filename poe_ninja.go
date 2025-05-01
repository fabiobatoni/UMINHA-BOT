package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Currency struct {
	CurrencyTypeName string  `json:"currencyTypeName"`
	ChaosValue       float64 `json:"chaosEquivalent"`
}

type Response struct {
	Lines []Currency `json:"lines"`
}

var cache []Currency
var lastUpdate time.Time
var mu sync.Mutex

func fetchData() []Currency {
	resp, err := http.Get("https://poe.ninja/api/data/currencyoverview?league=Necropolis&type=Currency")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var data Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil
	}

	return data.Lines
}

func ensureCache() {
	mu.Lock()
	defer mu.Unlock()

	if time.Since(lastUpdate) > 5*time.Minute || cache == nil {
		cache = fetchData()
		lastUpdate = time.Now()
	}
}

func BuscarTopPromos() string {
	ensureCache()

	if cache == nil {
		return "Não consegui acessar os dados agora."
	}

	texto := "**Top moedas - Path of Exile**\n"
	for i, item := range cache {
		if i >= 5 {
			break
		}
		texto += fmt.Sprintf("- %s: %.2f chaos\n", item.CurrencyTypeName, item.ChaosValue)
	}
	return texto
}

func BuscarPrecoItem(nome string) string {
	ensureCache()

	nome = strings.ToLower(strings.TrimSpace(nome))

	for _, item := range cache {
		if strings.ToLower(item.CurrencyTypeName) == nome {
			return fmt.Sprintf("💰 %s custa %.2f chaos.", item.CurrencyTypeName, item.ChaosValue)
		}
	}

	return fmt.Sprintf("❌ Item '%s' não encontrado.", nome)
}
