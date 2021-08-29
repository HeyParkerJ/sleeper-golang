package api

type Transactions struct {
	Type          string        `json:"type"`
	TransactionID string        `json:"transaction_id"`
	StatusUpdated int64         `json:"status_updated"`
	Status        string        `json:"status"`
	Settings      interface{}   `json:"settings"`
	RosterIds     []int         `json:"roster_ids"`
	Metadata      interface{}   `json:"metadata"`
	Leg           int           `json:"leg"`
	Drops         struct{}      `json:"drops"`
	DraftPicks    []interface{} `json:"draft_picks"`
	Creator       string        `json:"creator"`
	Created       int64         `json:"created"`
	ConsenterIds  []int         `json:"consenter_ids"`
	Adds          struct{}      `json:"adds"`
	WaiverBudget  []interface{} `json:"waiver_budget"`
}
