package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/llucasshenrique/mtgmind/definitions"
	"gorm.io/gorm/clause"
)

func PrepareDatabase(source string, database string, force bool) {
	var preparedCards []Card
	cardsArray := OpenJson(source)
	connection := CreateConnection(database)
	connection.AutoMigrate(&Card{})
	for _, card := range cardsArray {
		pCard := transformIntoCard(card)
		preparedCards = append(preparedCards, pCard)
	}
	if !force {
		connection.CreateInBatches(&preparedCards, 1000)
	} else {
		connection.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(&preparedCards, 1000)
	}
}

func OpenJson(path string) []definitions.JsonCardData {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	var payload []definitions.JsonCardData
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error unmarshalling json: ", err)
	}
	return payload
}

func transformIntoCard(card definitions.JsonCardData) Card {
	return Card{
		ID:            card.ID,
		Name:          card.Name,
		Type:          card.TypeLine,
		ManaCost:      card.ManaCost,
		Cmc:           card.Cmc,
		Power:         card.Power,
		Toughness:     card.Toughness,
		Colors:        strings.Join(card.Colors, ","),
		ColorIdentity: strings.Join(card.ColorIdentity, ","),
		Keywords:      strings.Join(card.Keywords, ","),
		Rarity:        card.Rarity,
	}
}
