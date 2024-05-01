package app

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	connStr := "postgres://postgres:admin@localhost/db_cat_social?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	var count int
	db.QueryRow("SELECT count(*) FROM user").Scan(&count)
	// Mengonversi nilai count menjadi string
	countStr := fmt.Sprintf("%d", count)

	// Mencetak nilai count dalam bentuk string
	fmt.Println("Jumlah baris:", countStr)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	fmt.Println("Connected to database")
	return db
}
