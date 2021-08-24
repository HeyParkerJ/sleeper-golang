package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Verification   interface{} `json:"verification"`
	Username       string      `json:"username"`
	UserID         string      `json:"user_id"`
	Token          interface{} `json:"token"`
	SummonerRegion interface{} `json:"summoner_region"`
	SummonerName   interface{} `json:"summoner_name"`
	Solicitable    interface{} `json:"solicitable"`
	RealName       interface{} `json:"real_name"`
	Phone          interface{} `json:"phone"`
	Pending        interface{} `json:"pending"`
	Notifications  interface{} `json:"notifications"`
	IsBot          bool        `json:"is_bot"`
	Email          interface{} `json:"email"`
	DisplayName    string      `json:"display_name"`
	Deleted        interface{} `json:"deleted"`
	DataUpdated    interface{} `json:"data_updated"`
	Currencies     interface{} `json:"currencies"`
	Created        interface{} `json:"created"`
	Cookies        interface{} `json:"cookies"`
	Avatar         string      `json:"avatar"`
}

type League struct {
	TotalRosters          int    `json:"total_rosters"`
	Status                string `json:"status"`
	Sport                 string `json:"sport"`
	Shard                 int    `json:"shard"`
	Settings              interface{}
	SeasonType            string      `json:"season_type"`
	Season                string      `json:"season"`
	ScoringSettings       interface{} `json:"scoring_settings"`
	RosterPositions       []string    `json:"roster_positions"`
	PreviousLeagueID      string      `json:"previous_league_id"`
	Name                  string      `json:"name"`
	Metadata              interface{} `json:"metadata"`
	LoserBracketID        interface{} `json:"loser_bracket_id"`
	LeagueID              string      `json:"league_id"`
	LastTransactionID     int64       `json:"last_transaction_id"`
	LastReadID            string      `json:"last_read_id"`
	LastPinnedMessageID   interface{} `json:"last_pinned_message_id"`
	LastMessageTime       int64       `json:"last_message_time"`
	LastMessageTextMap    interface{} `json:"last_message_text_map"`
	LastMessageID         string      `json:"last_message_id"`
	LastMessageAttachment interface{} `json:"last_message_attachment"`
	LastAuthorIsBot       bool        `json:"last_author_is_bot"`
	LastAuthorID          string      `json:"last_author_id"`
	LastAuthorDisplayName string      `json:"last_author_display_name"`
	LastAuthorAvatar      interface{} `json:"last_author_avatar"`
	GroupID               interface{} `json:"group_id"`
	DraftID               string      `json:"draft_id"`
	DisplayOrder          int         `json:"display_order"`
	CompanyID             interface{} `json:"company_id"`
	BracketID             interface{} `json:"bracket_id"`
	Avatar                string      `json:"avatar"`
}

func main() {
	GetUser()
	GetLeague()
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

const (
	baseURL = "https://api.sleeper.app/v1/user/445264202878152704"
)

func newClient(url string) (body []byte) {
	newURL := baseURL + url
	resp, err := http.Get(newURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

func GetUser() {
	body := newClient("")
	var user User
	if err := json.Unmarshal([]byte(body), &user); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(PrettyPrint(user))
}

func GetLeague() {
	body := newClient("/leagues/nfl/2021")
	var result []League
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	for _, v := range result {
		fmt.Println(PrettyPrint(v))
	}
}
