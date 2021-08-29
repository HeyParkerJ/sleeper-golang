package api

type Roster struct {
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
