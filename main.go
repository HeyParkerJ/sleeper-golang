package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	GetUser()
	GetLeague()
	GetRoster()
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

const (
	baseURL  = "https://api.sleeper.app/v1/"
	leagueID = "league/654061850773790720"
	userID   = "user/445264202878152704"
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
	body := newClient(userID)
	var user User
	if err := json.Unmarshal([]byte(body), &user); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(PrettyPrint(user))
}

func GetLeague() {
	// league url via docs https://docs.sleeper.app/#leagues
	body := newClient(leagueID)
	var result League
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(PrettyPrint(result))
}

func GetRoster() {
	// roster url via docs https://docs.sleeper.app/#getting-rosters-in-a-league
	rosterURL := leagueID + "/rosters"
	body := newClient(rosterURL)
	var result []Roster
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(PrettyPrint(result))
}

func GetUsers() {
	// users url via docs https://docs.sleeper.app/#getting-users-in-a-league
	userURL := leagueID + "/users"
	body := newClient(userURL)
	var result []Roster
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(PrettyPrint(result))
}
