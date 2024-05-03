package app

import (
	"cat-social-be/helper"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	username := helper.Getenv("DATABASE_USERNAME", "root")
	password := helper.Getenv("DATABASE_PASSWORD", "")
	host := helper.Getenv("DATABASE_HOST", "127.0.0.1")
	port := helper.Getenv("DATABASE_PORT", "5432")
	database := helper.Getenv("DATABASE_NAME", "db_cat_social")
	fmt.Println(username, password, host, port, database)
	// connStr := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"
	// connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, port, database)
	connStr := "postgres://postgres:admin@localhost:5432/db_cat_social?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	fmt.Println("Connected to database")
	return db
}
