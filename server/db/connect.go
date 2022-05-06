package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	POSTGRES_HOST     = "db"
	POSTGRES_PORT     = "5432"
	POSTGRES_USER     = "user"
	POSTGRES_PASSWORD = "password"
	POSTGRES_DBNAME   = "go-graphql-sandbox"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	opt := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST,
		POSTGRES_PORT,
		POSTGRES_USER,
		POSTGRES_PASSWORD,
		POSTGRES_DBNAME,
	)
	DB, err = gorm.Open(
		"postgres",
		opt,
	)
	if err != nil {
		log.Fatal(err)
	}
}
