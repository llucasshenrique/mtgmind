package search

import (
	"log"

	"github.com/llucasshenrique/mtgmind/database"
	"github.com/llucasshenrique/mtgmind/models"
)

func GetCardBySet(set string, databasePath string) {
	var cards [][]string
	result := searchCardsBySet(set, databasePath)
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

func searchCardsBySet(set string, databasePath string) []models.Card {
	connection := database.CreateConnection(databasePath)
	var searchedCard = models.Card{Set: set}
	var cards []models.Card
	search := connection.Where(&searchedCard).Find(&cards)
	if search.Error != nil {
		log.Fatal(search.Error)
	}
	return cards
}
