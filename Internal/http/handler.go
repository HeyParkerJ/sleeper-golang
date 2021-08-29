package handler

import (
	"encoding/json"
	"fmt"
	"github.com/MasEo9/sleeper-golang/Internal/api"
	"github.com/MasEo9/sleeper-golang/tools"
	"io"
	"log"
	"net/http"
	"strconv"
)

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
	var user api.User
	if err := json.Unmarshal([]byte(body), &user); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(user))
}

func GetLeague() {
	// league url via docs https://docs.sleeper.app/#leagues
	body := newClient(leagueID)
	var result api.League
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}

func GetRoster() {
	// roster url via docs https://docs.sleeper.app/#getting-rosters-in-a-league
	rosterURL := leagueID + "/rosters"
	body := newClient(rosterURL)
	var result []api.Roster
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}

func GetUsers() {
	// users url via docs https://docs.sleeper.app/#getting-users-in-a-league
	userURL := leagueID + "/users"
	body := newClient(userURL)
	var result []api.Roster
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}

func GetMatchups() {
	// users url via docs https://docs.sleeper.app/#getting-matchups-in-a-league
	for i := 1; i < 30; i++ {
		i := strconv.Itoa(i)
		matchupsURL := leagueID + "/matchups/" + i
		body := newClient(matchupsURL)
		var result []api.Matchups
		if err := json.Unmarshal([]byte(body), &result); err != nil {
			log.Fatalln(err)
		}
		if len(result) == 0 {
			break
		}
		fmt.Println(tools.PrettyPrint(result))
	}
}

func GetTransactions() {
	// users url via docs https://docs.sleeper.app/#getting-matchups-in-a-league
	for i := 1; i >= 1; i++ {
		i := strconv.Itoa(i)
		transactionsURL := leagueID + "/transactions/" + i
		body := newClient(transactionsURL)
		var result []api.Transactions
		if err := json.Unmarshal([]byte(body), &result); err != nil {
			log.Fatalln(err)
		}
		if len(result) == 0 {
			break
		}
		fmt.Println(tools.PrettyPrint(result))
	}
}

func GetTradedPicks() {
	// roster url via docs https://docs.sleeper.app/#get-traded-picks
	tradedPicksURL := leagueID + "/traded_picks"
	body := newClient(tradedPicksURL)
	var result []api.TradedPicks
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}

func GetNflState() {
	// roster url via docs https://docs.sleeper.app/#get-nfl-state
	nflStateURL := "state/nfl"
	body := newClient(nflStateURL)
	var result api.NflState
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}
