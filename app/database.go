package app

import (
	"database/sql"
	"time"
)

func NewConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/v_product")
	if err != nil {
		panic("Cannot Create Connection " + err.Error()) // 500 Internal Server Error
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
