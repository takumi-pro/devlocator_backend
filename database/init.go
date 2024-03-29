package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func DBConnect() (*sql.DB, error) {
	// TODO: 環境変数から取得
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_MAGE_PORT")
	dbHost := os.Getenv("DB_MAGE_HOST")
	dbTls := os.Getenv("DB_TLS")
	dbConn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=%s&interpolateParams=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		dbTls,
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
