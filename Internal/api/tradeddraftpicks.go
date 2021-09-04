package api

type TradedDraftPicks struct {
	Season          string `json:"season"`
	Round           int    `json:"round"`
	RosterID        int    `json:"roster_id"`
	PreviousOwnerID int    `json:"previous_owner_id"`
	OwnerID         int    `json:"owner_id"`
	DraftID         int64  `json:"draft_id"`
}
