package api

type DraftPicks struct {
	PlayerID          string        `json:"player_id"`
	PickedBy          string        `json:"picked_by"`
	RosterID          int           `json:"roster_id,omitempty"`
	Round             int           `json:"round,omitempty"`
	DraftSlot         int           `json:"draft_slot,omitempty"`
	PickNo            int           `json:"pick_no"`
	DraftPickMetadata DraftPickMeta `json:"metadata"`
	IsKeeper          interface{}   `json:"is_keeper"`
	DraftID           string        `json:"draft_id"`
}

type DraftPickMeta struct {
	Team         string `json:"team"`
	Status       string `json:"status"`
	Sport        string `json:"sport"`
	Position     string `json:"position"`
	PlayerID     string `json:"player_id"`
	Number       string `json:"number"`
	NewsUpdated  string `json:"news_updated"`
	LastName     string `json:"last_name"`
	InjuryStatus string `json:"injury_status"`
	FirstName    string `json:"first_name"`
}
