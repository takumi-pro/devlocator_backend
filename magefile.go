//go:build mage

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

type ConnpassApi struct {
	ResultsStart     int     `json:"results_start"`
	ResultsReturned  int     `json:"results_returned"`
	ResultsAvailable int     `json:"results_available"`
	Events           []Event `json:"events"`
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

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// バッチ処理のテスト
func BatchTest() error {
	today := time.Now()
	testDate := time.Date(2024, 3, 1, 0, 0, 0, 0, today.Location())

	dates := getDatesUntilNextMonth(testDate)
	fmt.Println(dates)
	return nil
}

// 引数に指定した日付から1ヶ月後の日付までの日付の配列を返却する
func getDatesUntilNextMonth(today time.Time) []string {
	currentYear, currentMonth, currentDay := today.Date()

	// 翌月の最終日を求める（現在が12とする）
	nextMonthTime := time.Date(currentYear, currentMonth+2, 1, 0, 0, 0, 0, today.Location()).AddDate(0, 0, -1)
	nextYear := nextMonthTime.Year()
	nextMonth := nextMonthTime.Month()

	// 今日の日にちと翌月の最終日を比較して
	// 翌月の方が大きい場合には終了日は今日の日にちと同じ値
	// 今日の日にちの方が大きい場合には翌月の最終日が終了日となる
	endDay := nextMonthTime.Day()
	if nextMonthTime.Day() > currentDay {
		endDay = currentDay
	}

	var dates []string
	month := time.Date(nextYear, nextMonth, endDay, 0, 0, 0, 0, today.Location())
	for !today.After(month) {
		dates = append(dates, today.Format("20060102"))
		today = today.AddDate(0, 0, 1)
	}
	return dates
}

// DBに接続できるかテストする
func dBConnect() (*sql.DB, error) {
	dbUser := "takumi"
	dbPassword := "password"
	dbName := "devlocator"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbName)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Printf("database connection error: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("ping err: %v", err)
		return nil, err
	}

	return db, nil
}

func getEvents(count int) (ConnpassApi, error) {
	url := fmt.Sprintf("https://connpass.com/api/v1/event?count=%d", count)
	res, err := http.Get(url)
	if err != nil {
		return ConnpassApi{}, fmt.Errorf("api request failed: %v", err)
	}
	defer res.Body.Close()

	var connpassApi ConnpassApi
	if err := json.NewDecoder(res.Body).Decode(&connpassApi); err != nil {
		return ConnpassApi{}, fmt.Errorf("response decode failed: %v", err)
	}

	return connpassApi, nil
}

func insertEvent(db *sql.DB, event Event) error {
	sqlStr := `
		INSERT INTO events (
			event_id,
			title,
			description
		) values (?, ?, ?);
	`

	newEvent := Event{
		EventId:     1,
		Title:       "test",
		Description: "test description",
	}

	_, err := db.Exec(sqlStr, newEvent.EventId, newEvent.Title, newEvent.Description)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// connpass apiからイベント情報の取得
func EventDisplay() error {
	connpassApi, err := getEvents(1)
	if err != nil {
		fmt.Println(err)
		return err
	}

	db, err := dBConnect()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	insertEvent(db, connpassApi.Events[0])

	fmt.Printf("event: %s", connpassApi.Events[0].Title)
	return nil
}

// ================= ↓default task↓ ===================

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "main", ".")
	return cmd.Run()
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./MyApp", "/usr/bin/MyApp")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}
