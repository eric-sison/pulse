package internal

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	connStr := os.Getenv("DB_URL")

	if connStr == "" {
		log.Fatal("Failed to connect to database. No connection string found.")
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	// db.SetMaxIdleConns(utils.ToInt(os.Getenv("DB_MAX_OPEN_CONN")))

	// db.SetMaxIdleConns(utils.ToInt(os.Getenv("DB_MAX_IDLE_CONN")))

	// db.SetConnMaxLifetime(time.Minute * time.Duration(utils.ToInt(os.Getenv("DB_MAX_LIFETIME"))))

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	return db, nil
}
