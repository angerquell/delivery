package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"delivery/config"
)

var DB *sql.DB 

func InitDB(cfg config.Config) error {
	var err error
	DB, err = sql.Open("postgres", cfg.DSN())
	if err != nil {
		return fmt.Errorf("failed to open a DB connection: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to the DB: %w", err)
	}

	fmt.Println("Successfully connected!")
	return nil
}
