package api

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
