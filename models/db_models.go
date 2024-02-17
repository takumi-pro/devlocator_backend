package models

type DBEventsResponse struct {
	ResultsReturned int64     `json:"results_returned"`
	Events          []DBEvent `json:"events"`
}

type DBEvent struct {
	EventId          int    `json:"event_id" gorm:"primaryKey"`
	Title            string `json:"title"`
	Catch            string `json:"catch,omitempty"`
	Description      string `json:"description,omitempty"`
	EventUrl         string `json:"event_url"`
	StartedAt        string `json:"started_at"`
	EndedAt          string `json:"ended_at"`
	Limit            int    `json:"limit"`
	HashTag          string `json:"hash_tag,omitempty"`
	EventType        string `json:"event_type,omitempty"`
	Accepted         int    `json:"accepted"`
	Waiting          int    `json:"waiting"`
	UpdatedAt        string `json:"updated_at"`
	OwnerId          int    `json:"owner_id,omitempty"`
	OwnerNickname    string `json:"owner_nickname,omitempty"`
	OwnerDisplayName string `json:"owner_display_name,omitempty"`
	Place            string `json:"place"`
	Address          string `json:"address"`
	Lat              string `json:"lat"`
	Lon              string `json:"lon"`
}

type DBSeries struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}
