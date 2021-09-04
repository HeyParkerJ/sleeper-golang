package api

type Player struct {
	PlayerData struct {
		Hashtag               string      `json:"hashtag"`
		DepthChartPosition    int         `json:"depth_chart_position"`
		Status                string      `json:"status"`
		Sport                 string      `json:"sport"`
		FantasyPositions      []string    `json:"fantasy_positions"`
		Number                int         `json:"number"`
		SearchLastName        string      `json:"search_last_name"`
		InjuryStartDate       interface{} `json:"injury_start_date"`
		Weight                string      `json:"weight"`
		Position              string      `json:"position"`
		PracticeParticipation interface{} `json:"practice_participation"`
		SportradarID          string      `json:"sportradar_id"`
		Team                  string      `json:"team"`
		LastName              string      `json:"last_name"`
		College               string      `json:"college"`
		FantasyDataID         int         `json:"fantasy_data_id"`
		InjuryStatus          interface{} `json:"injury_status"`
		PlayerID              string      `json:"player_id"`
		Height                string      `json:"height"`
		SearchFullName        string      `json:"search_full_name"`
		Age                   int         `json:"age"`
		StatsID               string      `json:"stats_id"`
		BirthCountry          string      `json:"birth_country"`
		EspnID                string      `json:"espn_id"`
		SearchRank            int         `json:"search_rank"`
		FirstName             string      `json:"first_name"`
		DepthChartOrder       int         `json:"depth_chart_order"`
		YearsExp              int         `json:"years_exp"`
		RotowireID            interface{} `json:"rotowire_id"`
		RotoworldID           int         `json:"rotoworld_id"`
		SearchFirstName       string      `json:"search_first_name"`
		YahooID               interface{} `json:"yahoo_id"`
	} `json:"omitempty"`
}
