package api

type Matchups struct {
	StartersPoints []float64   `json:"starters_points"`
	Starters       []string    `json:"starters"`
	RosterID       int         `json:"roster_id"`
	Points         float64     `json:"points"`
	Players        []string    `json:"players"`
	MatchupID      int         `json:"matchup_id"`
	CustomPoints   interface{} `json:"custom_points"`
}
