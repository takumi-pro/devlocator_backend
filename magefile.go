//go:build mage

package main

import (
	"database/sql"
	"devlocator/database"
	"devlocator/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

const (
	MAX_EVENTS_PER_REQUEST   = 10
	EVENT_INDEX_INCREMENT    = 10
	REQUEST_INTERVAL_SECONDS = 6
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// 日付処理のテスト
func DatesTest() error {
	today := time.Now()
	testDate := time.Date(2023, 12, 17, 0, 0, 0, 0, today.Location())

	dates := getDatesUntilNextMonth(testDate)
	fmt.Println(dates)
	return nil
}

func DBConnectGormTest() error {
	_, err := database.DBConnectGorm()
	if err != nil {
		return err
	}
	return nil
}

func DBConnectTest() error {
	db, err := database.DBConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return err
	}

	fmt.Println("db connect")
	return nil
}

// イベント情報登録バッチテスト
func EventsBatchTest() error {
	// dates := getDatesUntilNextMonth(time.Now())
	// stringDates := strings.Join(dates, ",")

	db, err := database.DBConnect()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	if err := deleteEventsTransaction(db); err != nil {
		fmt.Println(err.Error())
		return err
	}

	beforeRequestTime := time.Now()
	startIndex := 1
	for {
		d := (time.Now()).Sub(beforeRequestTime).Seconds()
		fmt.Printf("time duration: %v\n", d)
		eventsResponse, err := getEvents("20240104", startIndex, MAX_EVENTS_PER_REQUEST)
		if err != nil {
			return err
		}

		if len(eventsResponse.Events) == 0 {
			break
		}

		if err := insertEventsTransaction(db, eventsResponse.Events); err != nil {
			return err
		}

		startIndex += EVENT_INDEX_INCREMENT
		time.Sleep(REQUEST_INTERVAL_SECONDS * time.Second)
	}

	return nil
}

// 引数に指定した日付から1ヶ月後の日付までの日付の配列を返却する
func getDatesUntilNextMonth(today time.Time) []string {
	currentYear, currentMonth, currentDay := today.Date()

	// 翌月の最終日を求める
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

// 引数
// count - 取得件数
// date - イベント日付
// start - 検索開始位置
func getEvents(date string, startIndex int, count int) (models.EventsResponse, error) {
	url := fmt.Sprintf("https://connpass.com/api/v1/event?count=%d&start=%d&ymd=%s", count, startIndex, date)
	res, err := http.Get(url)
	if err != nil {
		return models.EventsResponse{}, err
	}
	defer res.Body.Close()

	var connpassApi models.EventsResponse
	if err := json.NewDecoder(res.Body).Decode(&connpassApi); err != nil {
		return models.EventsResponse{}, err
	}

	// オフラインのみのイベント情報を取得するため
	// 経度と緯度を判定
	var filteredEvents []models.Event
	for _, event := range connpassApi.Events {
		if event.Lat != "" && event.Lon != "" {
			filteredEvents = append(filteredEvents, event)
		}
	}
	connpassApi.Events = filteredEvents

	return connpassApi, nil
}

func deleteEventsTransaction(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteEvents(tx); err != nil {
		return err
	}

	fmt.Println("deleted!")
	return tx.Commit()
}

func deleteEvents(tx *sql.Tx) error {
	sqlStr := `DELETE FROM events;`
	if _, err := tx.Exec(sqlStr); err != nil {
		return err
	}
	return nil
}

func insertEventsTransaction(db *sql.DB, events []models.Event) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := insertEvents(tx, events); err != nil {
		return err
	}

	return tx.Commit()
}

func insertEvents(tx *sql.Tx, events []models.Event) error {
	sqlStr := `
		INSERT INTO events (
			event_id,
			title,
			catch,
			description,
			event_url,
			started_at,
			ended_at,
			` + "`limit`" + `,
			hash_tag,
			event_type,
			accepted,
			waiting,
			owner_id,
			owner_nickname,
			owner_display_name,
			place,
			address,
			lat,
			lon
		) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, e := range events {
		if _, err := stmt.Exec(
			e.EventId, e.Title, e.Catch, e.Description, e.EventUrl, e.StartedAt,
			e.EndedAt, e.Limit, e.HashTag, e.EventType, e.Accepted, e.Waiting,
			e.OwnerId, e.OwnerNickname, e.OwnerDisplayName, e.Place, e.Address, e.Lat, e.Lon,
		); err != nil {
			return err
		}
	}

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
