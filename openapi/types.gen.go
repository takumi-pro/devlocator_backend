// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package openapi

import (
	"time"
)

const (
	BearerScopes = "bearer.Scopes"
)

// Error エラーモデル
type Error struct {
	// Code ステータスコード
	Code int `json:"code"`

	// Details エラーメッセージの詳細
	Details string `json:"details"`

	// Message エラーメッセージ
	Message string `json:"message"`
}

// Event イベントモデル
type Event struct {
	// Accepted 参加者数
	Accepted *int `json:"accepted,omitempty"`

	// Address 開催場所
	Address *string `json:"address,omitempty"`

	// Catch キャッチ
	Catch *string `json:"catch,omitempty"`

	// Description 概要
	Description *string `json:"description,omitempty"`

	// EndedAt イベント終了日時
	EndedAt *time.Time `json:"endedAt,omitempty"`

	// EventId イベントID
	EventId int `json:"eventId"`

	// EventType イベント参加タイプ
	// participation: connpassで参加受付あり
	// advertisement: 告知のみ
	EventType *string `json:"eventType,omitempty"`

	// EventUrl イベントURL
	EventUrl *string `json:"eventUrl,omitempty"`

	// Lat 開催会場の緯度
	Lat string `json:"lat"`

	// Limit 定員
	Limit *int `json:"limit,omitempty"`

	// Lon 開催会場の経度
	Lon string `json:"lon"`

	// Place 開催会場
	Place *string `json:"place,omitempty"`

	// StartedAt イベント開催日時
	StartedAt *time.Time `json:"startedAt,omitempty"`

	// Title イベントタイトル
	Title string `json:"title"`

	// UpdatedAt イベント更新日時
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// Waiting 補欠者数
	Waiting *int `json:"waiting,omitempty"`
}

// Date defines model for date.
type Date = string

// EventId defines model for event_id.
type EventId = string

// Keyword defines model for keyword.
type Keyword = string

// Prefecture defines model for prefecture.
type Prefecture = string

// SearchMethod defines model for search_method.
type SearchMethod = string

// N400BadRequest エラーモデル
type N400BadRequest = Error

// N401Unauthorized エラーモデル
type N401Unauthorized = Error

// N403Forbidden エラーモデル
type N403Forbidden = Error

// N404NotFound エラーモデル
type N404NotFound = Error

// N500InternalServerError エラーモデル
type N500InternalServerError = Error

// DetailEvent defines model for detail_event.
type DetailEvent struct {
	// Events イベント情報配列
	Events          []Event `json:"events"`
	ResultsReturned int     `json:"resultsReturned"`
}

// SearchEvent defines model for search_event.
type SearchEvent struct {
	// Events イベントリスト
	Events []struct {
		Accepted  *int    `json:"accepted,omitempty"`
		Address   *string `json:"address,omitempty"`
		EndedAt   *string `json:"endedAt,omitempty"`
		EventId   int     `json:"eventId"`
		EventType *string `json:"eventType,omitempty"`
		EventUrl  *string `json:"eventUrl,omitempty"`
		Lat       string  `json:"lat"`
		Limit     *int    `json:"limit,omitempty"`
		Lon       string  `json:"lon"`
		Place     *string `json:"place,omitempty"`
		StartedAt *string `json:"startedAt,omitempty"`
		Title     string  `json:"title"`
		UpdatedAt *string `json:"updatedAt,omitempty"`
		Waiting   *int    `json:"waiting,omitempty"`
	} `json:"events"`

	// ResultsReturned 含まれる検索結果の件数
	ResultsReturned int `json:"resultsReturned"`
}

// Users defines model for users.
type Users struct {
	FirebaseUid string `json:"firebaseUid"`

	// Image googleアカウントのアイコン画像
	Image *string `json:"image,omitempty"`

	// MarkedEvents ブックマークしたイベント
	MarkedEvents []Event `json:"markedEvents"`

	// Name 名前
	Name string `json:"name"`
}

// Bookmark defines model for bookmark.
type Bookmark struct {
	// Id eventsテーブルのid
	Id string `json:"id"`
}

// GetApiEventParams defines parameters for GetApiEvent.
type GetApiEventParams struct {
	// EventId イベント毎に割り当てられた番号
	EventId *EventId `form:"event_id,omitempty" json:"event_id,omitempty"`

	// Keyword キーワード
	Keyword *Keyword `form:"keyword,omitempty" json:"keyword,omitempty"`

	// SearchMethod 検索条件（ORもしくはAND）
	SearchMethod *SearchMethod `form:"search_method,omitempty" json:"search_method,omitempty"`

	// Date イベント開催日
	Date *Date `form:"date,omitempty" json:"date,omitempty"`

	// Prefecture イベント開催都道府県
	Prefecture *Prefecture `form:"prefecture,omitempty" json:"prefecture,omitempty"`
}

// PutApiEventBookmarkJSONBody defines parameters for PutApiEventBookmark.
type PutApiEventBookmarkJSONBody struct {
	// Id eventsテーブルのid
	Id string `json:"id"`
}

// PutApiEventBookmarkJSONRequestBody defines body for PutApiEventBookmark for application/json ContentType.
type PutApiEventBookmarkJSONRequestBody PutApiEventBookmarkJSONBody
