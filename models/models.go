package models

type EventsResponse struct {
	ResultsReturned int     `json:"results_returned"`
	Events          []Event `json:"events"`
}

type Event struct {
	EventId          int    `json:"event_id"`
	Title            string `json:"title"`
	Catch            string `json:"catch"`
	Description      string `json:"description"`
	EventUrl         string `json:"event_url"`
	StartedAt        string `json:"started_at"`
	EndedAt          string `json:"ended_at"`
	Limit            int    `json:"limit"`
	HashTag          string `json:"hash_tag"`
	EventType        string `json:"event_target"`
	Accepted         int    `json:"accepted"`
	Waiting          int    `json:"waiting"`
	UpdatedAt        string `json:"updated_at"`
	OwnerId          int    `json:"owner_id"`
	OwnerNickname    string `json:"owner_nickname"`
	OwnerDisplayName string `json:"owner_display_name"`
	Place            string `json:"place"`
	Address          string `json:"address"`
	Lat              string `json:"lat"`
	Lon              string `json:"lon"`
	Series           Series `json:"series"`
}

type Series struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}
