package api

type LeagueDraft struct {
	Type            string        `json:"type"`
	Status          string        `json:"status"`
	StartTime       int64         `json:"start_time"`
	Sport           string        `json:"sport"`
	DraftSettings   DraftSettings `json:"settings"`
	SeasonType      string        `json:"season_type"`
	Season          string        `json:"season"`
	DraftMetadata   DraftMeta     `json:"metadata"`
	LeagueID        string        `json:"league_id"`
	LastPicked      int64         `json:"last_picked"`
	LastMessageTime int64         `json:"last_message_time"`
	LastMessageID   string        `json:"last_message_id"`
	DraftOrder      interface{}   `json:"draft_order"`
	DraftID         string        `json:"draft_id"`
	Creators        interface{}   `json:"creators"`
	Created         int64         `json:"created"`
}

type DraftMeta struct {
	ScoringType string `json:"scoring_type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DraftSettings struct {
	Teams     int `json:"teams"`
	SlotsWr   int `json:"slots_wr"`
	SlotsTe   int `json:"slots_te"`
	SlotsRb   int `json:"slots_rb"`
	SlotsQb   int `json:"slots_qb"`
	SlotsK    int `json:"slots_k"`
	SlotsFlex int `json:"slots_flex"`
	SlotsDef  int `json:"slots_def"`
	SlotsBn   int `json:"slots_bn"`
	Rounds    int `json:"rounds"`
	PickTimer int `json:"pick_timer"`
}
