package database

import (
	"database/sql"
	"fmt"
)

func DBConnect() (*sql.DB, error) {
	// TODO: 環境変数から取得
	dbUser := "takumi"
	dbPassword := "password"
	dbName := "devlocator"
	dbPort := "3307"
	dbHost := "127.0.0.1"
	dbConn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
