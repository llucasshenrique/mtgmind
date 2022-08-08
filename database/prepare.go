package database

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/llucasshenrique/mtgmind/models"
	"gorm.io/gorm/clause"
)

func PrepareDatabase(source string, database string, force bool) {
	var preparedCards []models.Card
	cardsArray := OpenJson(source)
	connection := CreateConnection(database)
	connection.AutoMigrate(&models.Card{})
	for _, card := range cardsArray {
		preparedCards = append(preparedCards, card.ToCard())
	}
	if !force {
		connection.CreateInBatches(&preparedCards, 1000)
	} else {
		connection.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(&preparedCards, 1000)
	}
}

func OpenJson(path string) []JsonCardData {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	var payload []JsonCardData
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error unmarshalling json: ", err)
	}
	return payload
}
