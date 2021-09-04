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
	// matchups url via docs https://docs.sleeper.app/#getting-matchups-in-a-league
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
	// transactions url via docs https://docs.sleeper.app/#getting-matchups-in-a-league
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
	// traded picks url via docs https://docs.sleeper.app/#get-traded-picks
	tradedPicksURL := leagueID + "/traded_picks"
	body := newClient(tradedPicksURL)
	var result []api.TradedPicks
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}

func GetNflState() {
	// nfl current state url via docs https://docs.sleeper.app/#get-nfl-state
	nflStateURL := "state/nfl"
	body := newClient(nflStateURL)
	var result api.NflState
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}

func GetLeagueDraft() (draftID string) {
	// league draft url via docs https://docs.sleeper.app/#get-all-drafts-for-a-league
	leagueDraftURL := leagueID + "/drafts"
	body := newClient(leagueDraftURL)
	var result []api.LeagueDraft
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	for _, v := range result {
		draftID = v.DraftID
	}
	// fmt.Println(tools.PrettyPrint(result))
	return draftID
}

// get draft id from league draft pull
var draftID = GetLeagueDraft()

func GetDraftPicks() {
	// draft picks url via docs https://docs.sleeper.app/#get-all-picks-in-a-draft
	draftPicksURL := "draft/" + draftID + "/picks"
	body := newClient(draftPicksURL)
	var result []api.DraftPicks
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}

func GetTradedDraftPicks() {
	// traded draft picks url via docs https://docs.sleeper.app/#get-traded-picks-in-a-draft
	tradeddraftPicksURL := "draft/" + draftID + "/traded_picks"
	body := newClient(tradeddraftPicksURL)
	var result []api.TradedDraftPicks
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))

}

func GetPlayers() {
	// league url via docs https://docs.sleeper.app/#fetch-all-players
	playersURL := "players/nfl"
	body := newClient(playersURL)
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}

func GetStats() {
	// stats url via docs https://web.archive.org/web/20191208152134/https://docs.sleeper.app/#stats
	// GET https://api.sleeper.app/v1/stats/<sport>/<season_type>/<season>
	// GET https://api.sleeper.app/v1/stats/<sport>/<season_type>/<season>/<week>
	// GET https://api.sleeper.app/v1/projections/<sport>/<season_type>/<season>
	// GET https://api.sleeper.app/v1/projections/<sport>/<season_type>/<season>/<week>

	// url set for test data load with preseason week 3 data
	playersURL := "stats/nfl/pre/2021/3"
	body := newClient(playersURL)
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(body),
		&result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tools.PrettyPrint(result))
}
