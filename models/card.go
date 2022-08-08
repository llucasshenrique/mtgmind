package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	ID            string
	Name          string
	Type          string
	ManaCost      string
	Cmc           float64
	Power         string
	Toughness     string
	Colors        string
	ColorIdentity string
	Keywords      string
	Rarity        string
	Set           string
}
