package search

import (
	"log"

	"github.com/llucasshenrique/mtgmind/database"
)

func GetCardByName(name string, databasePath string) {
	var cards [][]string
	headers := []string{"ID", "Name", "Mana Cost", "Type"}
	result := searchCardsByName(name, databasePath)
	for _, card := range result {
		cards = append(cards, []string{card.ID, card.Name, card.ManaCost, card.Type})
	}
	RenderTable(headers, cards)
}

func searchCardsByName(name string, databasePath string) []database.Card {
	connection := database.CreateConnection(databasePath)
	var searchedCard = database.Card{Name: name}
	var cards []database.Card
	search := connection.Where(&searchedCard).Find(&cards)
	if search.Error != nil {
		log.Fatal(search.Error)
	}
	return cards
}
