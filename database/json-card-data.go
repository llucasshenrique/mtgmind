package database

import (
	"strings"

	"github.com/llucasshenrique/mtgmind/models"
)

type ImageUris struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	Png        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

type Prices struct {
	Usd       string      `json:"usd"`
	UsdFoil   string      `json:"usd_foil"`
	UsdEtched interface{} `json:"usd_etched"`
	Eur       string      `json:"eur"`
	EurFoil   string      `json:"eur_foil"`
	Tix       string      `json:"tix"`
}

type RelatedUris struct {
	Gatherer                  string `json:"gatherer"`
	TcgplayerInfiniteArticles string `json:"tcgplayer_infinite_articles"`
	TcgplayerInfiniteDecks    string `json:"tcgplayer_infinite_decks"`
	Edhrec                    string `json:"edhrec"`
}

type Legalities struct {
	Standard        string `json:"standard"`
	Future          string `json:"future"`
	Historic        string `json:"historic"`
	Gladiator       string `json:"gladiator"`
	Pioneer         string `json:"pioneer"`
	Explorer        string `json:"explorer"`
	Modern          string `json:"modern"`
	Legacy          string `json:"legacy"`
	Pauper          string `json:"pauper"`
	Vintage         string `json:"vintage"`
	Penny           string `json:"penny"`
	Commander       string `json:"commander"`
	Brawl           string `json:"brawl"`
	Historicbrawl   string `json:"historicbrawl"`
	Alchemy         string `json:"alchemy"`
	Paupercommander string `json:"paupercommander"`
	Duel            string `json:"duel"`
	Oldschool       string `json:"oldschool"`
	Premodern       string `json:"premodern"`
}

type JsonCardData struct {
	Object          string      `json:"object"`
	ID              string      `json:"id"`
	OracleID        string      `json:"oracle_id"`
	MultiverseIds   []int       `json:"multiverse_ids"`
	MtgoID          int         `json:"mtgo_id"`
	MtgoFoilID      int         `json:"mtgo_foil_id"`
	TcgplayerID     int         `json:"tcgplayer_id"`
	CardmarketID    int         `json:"cardmarket_id"`
	Name            string      `json:"name"`
	Lang            string      `json:"lang"`
	ReleasedAt      string      `json:"released_at"`
	URI             string      `json:"uri"`
	ScryfallURI     string      `json:"scryfall_uri"`
	Layout          string      `json:"layout"`
	HighresImage    bool        `json:"highres_image"`
	ImageStatus     string      `json:"image_status"`
	ImageUris       ImageUris   `json:"image_uris"`
	ManaCost        string      `json:"mana_cost"`
	Cmc             float64     `json:"cmc"`
	TypeLine        string      `json:"type_line"`
	OracleText      string      `json:"oracle_text"`
	Power           string      `json:"power"`
	Toughness       string      `json:"toughness"`
	Colors          []string    `json:"colors"`
	ColorIdentity   []string    `json:"color_identity"`
	Keywords        []string    `json:"keywords"`
	Legalities      Legalities  `json:"legalities"`
	Games           []string    `json:"games"`
	Reserved        bool        `json:"reserved"`
	Foil            bool        `json:"foil"`
	Nonfoil         bool        `json:"nonfoil"`
	Finishes        []string    `json:"finishes"`
	Oversized       bool        `json:"oversized"`
	Promo           bool        `json:"promo"`
	Reprint         bool        `json:"reprint"`
	Variation       bool        `json:"variation"`
	SetID           string      `json:"set_id"`
	Set             string      `json:"set"`
	SetName         string      `json:"set_name"`
	SetType         string      `json:"set_type"`
	SetURI          string      `json:"set_uri"`
	SetSearchURI    string      `json:"set_search_uri"`
	ScryfallSetURI  string      `json:"scryfall_set_uri"`
	RulingsURI      string      `json:"rulings_uri"`
	PrintsSearchURI string      `json:"prints_search_uri"`
	CollectorNumber string      `json:"collector_number"`
	Digital         bool        `json:"digital"`
	Rarity          string      `json:"rarity"`
	FlavorText      string      `json:"flavor_text"`
	CardBackID      string      `json:"card_back_id"`
	Artist          string      `json:"artist"`
	ArtistIds       []string    `json:"artist_ids"`
	IllustrationID  string      `json:"illustration_id"`
	BorderColor     string      `json:"border_color"`
	Frame           string      `json:"frame"`
	FullArt         bool        `json:"full_art"`
	Textless        bool        `json:"textless"`
	Booster         bool        `json:"booster"`
	StorySpotlight  bool        `json:"story_spotlight"`
	EdhrecRank      int         `json:"edhrec_rank"`
	PennyRank       int         `json:"penny_rank"`
	Prices          Prices      `json:"prices"`
	RelatedUris     RelatedUris `json:"related_uris"`
}

func (card JsonCardData) ToCard() models.Card {
	return models.Card{
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
		Set:           card.Set,
	}
}
