package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnectGorm() (db *gorm.DB, err error) {
	// TODO: 環境変数から取得
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
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

	db, err = gorm.Open(mysql.Open(dbConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected!")

	return db, nil
}
