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
	TotalRosters          int             `json:"total_rosters"`
	Status                string          `json:"status"`
	Sport                 string          `json:"sport"`
	Shard                 int             `json:"shard"`
	Settings              Settings        `json:"settings"`
	SeasonType            string          `json:"season_type"`
	Season                string          `json:"season"`
	ScoringSettings       ScoringSettings `json:"scoring_settings"`
	RosterPositions       []string        `json:"roster_positions"`
	PreviousLeagueID      string          `json:"previous_league_id"`
	Name                  string          `json:"name"`
	Metadata              interface{}     `json:"metadata"`
	LoserBracketID        interface{}     `json:"loser_bracket_id"`
	LeagueID              string          `json:"league_id"`
	LastReadID            interface{}     `json:"last_read_id"`
	LastPinnedMessageID   interface{}     `json:"last_pinned_message_id"`
	LastMessageTime       int64           `json:"last_message_time"`
	LastMessageTextMap    interface{}     `json:"last_message_text_map"`
	LastMessageID         string          `json:"last_message_id"`
	LastMessageAttachment interface{}     `json:"last_message_attachment"`
	LastAuthorIsBot       bool            `json:"last_author_is_bot"`
	LastAuthorID          string          `json:"last_author_id"`
	LastAuthorDisplayName string          `json:"last_author_display_name"`
	LastAuthorAvatar      interface{}     `json:"last_author_avatar"`
	GroupID               interface{}     `json:"group_id"`
	DraftID               string          `json:"draft_id"`
	CompanyID             interface{}     `json:"company_id"`
	BracketID             interface{}     `json:"bracket_id"`
	Avatar                string          `json:"avatar"`
}

type Settings struct {
	MaxKeepers           int `json:"max_keepers"`
	DraftRounds          int `json:"draft_rounds"`
	TradeReviewDays      int `json:"trade_review_days"`
	ReserveAllowDnr      int `json:"reserve_allow_dnr"`
	CapacityOverride     int `json:"capacity_override"`
	PickTrading          int `json:"pick_trading"`
	TaxiYears            int `json:"taxi_years"`
	TaxiAllowVets        int `json:"taxi_allow_vets"`
	LastReport           int `json:"last_report"`
	DisableAdds          int `json:"disable_adds"`
	WaiverType           int `json:"waiver_type"`
	BenchLock            int `json:"bench_lock"`
	ReserveAllowSus      int `json:"reserve_allow_sus"`
	Type                 int `json:"type"`
	ReserveAllowCov      int `json:"reserve_allow_cov"`
	WaiverClearDays      int `json:"waiver_clear_days"`
	DailyWaiversLastRan  int `json:"daily_waivers_last_ran"`
	WaiverDayOfWeek      int `json:"waiver_day_of_week"`
	StartWeek            int `json:"start_week"`
	PlayoffTeams         int `json:"playoff_teams"`
	NumTeams             int `json:"num_teams"`
	ReserveSlots         int `json:"reserve_slots"`
	PlayoffRoundType     int `json:"playoff_round_type"`
	DailyWaiversHour     int `json:"daily_waivers_hour"`
	WaiverBudget         int `json:"waiver_budget"`
	ReserveAllowOut      int `json:"reserve_allow_out"`
	OffseasonAdds        int `json:"offseason_adds"`
	PlayoffSeedType      int `json:"playoff_seed_type"`
	DailyWaivers         int `json:"daily_waivers"`
	Divisions            int `json:"divisions"`
	PlayoffWeekStart     int `json:"playoff_week_start"`
	DailyWaiversDays     int `json:"daily_waivers_days"`
	LeagueAverageMatch   int `json:"league_average_match"`
	Leg                  int `json:"leg"`
	TradeDeadline        int `json:"trade_deadline"`
	ReserveAllowDoubtful int `json:"reserve_allow_doubtful"`
	TaxiDeadline         int `json:"taxi_deadline"`
	ReserveAllowNa       int `json:"reserve_allow_na"`
	TaxiSlots            int `json:"taxi_slots"`
	PlayoffType          int `json:"playoff_type"`
}

type ScoringSettings struct {
	Pass2Pt      float64 `json:"pass_2pt"`
	PassInt      float64 `json:"pass_int"`
	Fgmiss       float64 `json:"fgmiss"`
	RecYd        float64 `json:"rec_yd"`
	IdpSafe      float64 `json:"idp_safe"`
	Xpmiss       float64 `json:"xpmiss"`
	DefPrTd      float64 `json:"def_pr_td"`
	Fgm3039      float64 `json:"fgm_30_39"`
	BlkKick      float64 `json:"blk_kick"`
	PtsAllow713  float64 `json:"pts_allow_7_13"`
	IdpFf        float64 `json:"idp_ff"`
	Ff           float64 `json:"ff"`
	Fgm2029      float64 `json:"fgm_20_29"`
	Fgm4049      float64 `json:"fgm_40_49"`
	PtsAllow16   float64 `json:"pts_allow_1_6"`
	StFumRec     float64 `json:"st_fum_rec"`
	DefStFf      float64 `json:"def_st_ff"`
	StFf         float64 `json:"st_ff"`
	IdpDefTd     float64 `json:"idp_def_td"`
	PtsAllow2834 float64 `json:"pts_allow_28_34"`
	IdpBlkKick   float64 `json:"idp_blk_kick"`
	IdpInt       float64 `json:"idp_int"`
	IdpQbHit     float64 `json:"idp_qb_hit"`
	Fgm50P       float64 `json:"fgm_50p"`
	FumRec       float64 `json:"fum_rec"`
	IdpTkl       float64 `json:"idp_tkl"`
	DefTd        float64 `json:"def_td"`
	Fgm019       float64 `json:"fgm_0_19"`
	IdpPassDef   float64 `json:"idp_pass_def"`
	IdpTklLoss   float64 `json:"idp_tkl_loss"`
	Int          float64 `json:"int"`
	PtsAllow0    float64 `json:"pts_allow_0"`
	PtsAllow2127 float64 `json:"pts_allow_21_27"`
	IdpFumRec    float64 `json:"idp_fum_rec"`
	Rec2Pt       float64 `json:"rec_2pt"`
	Rec          float64 `json:"rec"`
	Xpm          float64 `json:"xpm"`
	IdpSack      float64 `json:"idp_sack"`
	StTd         float64 `json:"st_td"`
	DefStFumRec  float64 `json:"def_st_fum_rec"`
	DefStTd      float64 `json:"def_st_td"`
	Sack         float64 `json:"sack"`
	IdpTklAst    float64 `json:"idp_tkl_ast"`
	FumRecTd     float64 `json:"fum_rec_td"`
	IdpTklSolo   float64 `json:"idp_tkl_solo"`
	Rush2Pt      float64 `json:"rush_2pt"`
	RecTd        float64 `json:"rec_td"`
	PtsAllow35P  float64 `json:"pts_allow_35p"`
	PtsAllow1420 float64 `json:"pts_allow_14_20"`
	RushYd       float64 `json:"rush_yd"`
	PassYd       float64 `json:"pass_yd"`
	PassTd       float64 `json:"pass_td"`
	RushTd       float64 `json:"rush_td"`
	DefKrTd      float64 `json:"def_kr_td"`
	FumLost      float64 `json:"fum_lost"`
	Fum          float64 `json:"fum"`
	Safe         float64 `json:"safe"`
}

type Roster []struct {
	Taxi      interface{}    `json:"taxi"`
	Starters  []string       `json:"starters"`
	Settings  RosterStats    `json:"settings"`
	RosterID  int            `json:"roster_id"`
	Reserve   interface{}    `json:"reserve"`
	Players   []string       `json:"players"`
	PlayerMap interface{}    `json:"player_map"`
	OwnerID   string         `json:"owner_id"`
	Metadata  RosterMetadata `json:"metadata,omitempty"`
}
type RosterStats struct {
	Wins             int `json:"wins"`
	WaiverPosition   int `json:"waiver_position"`
	WaiverBudgetUsed int `json:"waiver_budget_used"`
	TotalMoves       int `json:"total_moves"`
	Ties             int `json:"ties"`
	Losses           int `json:"losses"`
	Fpts             int `json:"fpts"`
	Division         int `json:"division"`
}

type RosterMetadata struct {
	Record string `json:"record"`
}

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
	// user url is the same as the base url
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
	var result Roster
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(PrettyPrint(result))
}
