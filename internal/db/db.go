package db

import (
	"cleangin/internal/constants"
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	gormdb, err := newDbClient(nil)

	if err != nil {
		panic("error while connecting database:" + err.Error())
	}

	return gormdb
}

func newDbClient(db *gorm.DB) (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	dsn := generateDns()
	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	return gormDb, nil
}

func generateDns() string {
	host := os.Getenv(constants.PostgresqlDBHost)
	port := os.Getenv(constants.PostgreqlDBPort)
	name := os.Getenv(constants.PostgresqlDBName)
	pass := os.Getenv(constants.PostgresqlDBPassword)
	user := os.Getenv(constants.PostgresqlDBUser)
	sslmode := os.Getenv(constants.PostgresqlDBSslMode)
	parts := []string{}
	if host != "" {
		parts = append(parts, fmt.Sprintf("host=%s", host))
	} else {
		parts = append(parts, fmt.Sprint("host=localhost"))
	}
	if port != "" {
		parts = append(parts, fmt.Sprintf("port=%s", port))
	} else {
		parts = append(parts, fmt.Sprint("port=5432"))
	}
	if name != "" {
		parts = append(parts, fmt.Sprintf("dbname=%s", name))
	} else {
		parts = append(parts, fmt.Sprint("dbname=postgres"))
	}
	if pass != "" {
		parts = append(parts, fmt.Sprintf("password=%s", pass))
	} else {
		parts = append(parts, fmt.Sprint("password=supersecretpassword"))
	}
	if user != "" {
		parts = append(parts, fmt.Sprintf("user=%s", user))
	} else {
		parts = append(parts, fmt.Sprint("user=postgres"))
	}
	if sslmode != "" {
		parts = append(parts, fmt.Sprintf("sslmode=%s", sslmode))
	} else {
		parts = append(parts, fmt.Sprint("sslmode=disable"))
	}
	return strings.Join(parts, " ")
}
