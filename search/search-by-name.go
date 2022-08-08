package search

import (
	"log"

	"github.com/llucasshenrique/mtgmind/database"
	"github.com/llucasshenrique/mtgmind/models"
)

func GetCardByName(name string, databasePath string) {
	var cards [][]string
	result := searchCardsByName(name, databasePath)
	headers := []string{"ID", "Name", "Mana Cost", "Type", "Set"}
	for _, card := range result {
		cards = append(cards, []string{
			card.ID,
			card.Name,
			card.ManaCost,
			card.Type,
			card.Set,
		})
	}
	RenderTable(headers, cards)
}

func searchCardsByName(name string, databasePath string) []models.Card {
	connection := database.CreateConnection(databasePath)
	var searchedCard = models.Card{Name: name}
	var cards []models.Card
	search := connection.Where(&searchedCard).Find(&cards)
	if search.Error != nil {
		log.Fatal(search.Error)
	}
	return cards
}
