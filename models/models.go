package models

type EventsResponse struct {
	ResultsReturned int64   `json:"resultsReturned"`
	Events          []Event `json:"events"`
}

type Event struct {
	EventId          int    `json:"eventId" gorm:"primaryKey"`
	Title            string `json:"title"`
	Catch            string `json:"catch,omitempty"`
	Description      string `json:"description,omitempty"`
	EventUrl         string `json:"eventUrl"`
	StartedAt        string `json:"startedAt"`
	EndedAt          string `json:"endedAt"`
	Limit            int    `json:"limit"`
	HashTag          string `json:"hashTag,omitempty"`
	EventType        string `json:"eventType,omitempty"`
	Accepted         int    `json:"accepted"`
	Waiting          int    `json:"waiting"`
	UpdatedAt        string `json:"updatedAt"`
	OwnerId          int    `json:"ownerId,omitempty"`
	OwnerNickname    string `json:"ownerNickname,omitempty"`
	OwnerDisplayName string `json:"ownerDisplayName,omitempty"`
	Place            string `json:"place"`
	Address          string `json:"address"`
	Lat              string `json:"lat"`
	Lon              string `json:"lon"`
}

type Series struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}
