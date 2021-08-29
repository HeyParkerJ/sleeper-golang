package api

type TradedPicks struct {
	Season          string `json:"season"`
	Round           int    `json:"round"`
	RosterID        int    `json:"roster_id"`
	PreviousOwnerID int    `json:"previous_owner_id"`
	OwnerID         int    `json:"owner_id"`
}
