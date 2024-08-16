package internal

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/eric-sison/pulse/pkg/utils"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	connStr := os.Getenv("DB_URL")

	// If the connection string is empty, it means that the necessary
	// information to connect to the database is missing or not provided. In this case, the code logs a
	// fatal error message and terminates the program using `log.Fatal`.
	if connStr == "" {
		log.Fatal("Failed to connect to database. No connection string found.")
	}

	// Attempt to open a new database connection using the postgres driver and the connection
	// string which contains the necessary information to connect to the database.
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	// Ensure that the database connection will be closed properly
	// after the function has finished executing, regardless of whether an error occurred or not.
	defer db.Close()

	setDbConfig(db)

	// This code snippet is checking the connection to the database by calling the `Ping()` method on the
	// database connection (`db`).
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	// Indicates that the database connection was successfully established without any errors,
	// allowing the calling code to use this connection for database operations.
	return db, nil
}

// Set the configuration parameters for a SQL database connection.
func setDbConfig(db *sql.DB) {

	maxIdleConns := os.Getenv("DB_MAX_OPEN_CONNS")

	maxOpenConns := os.Getenv("DB_MAX_OPEN_CONNS")

	maxLifetime := os.Getenv("DB_MAX_LIFETIME")

	maxIdleTime := os.Getenv("DB_MAX_IDLE_TIME")

	// Checking if any of the configuration parameters for the database connection
	// pool are empty. If any of the parameters (`maxIdleConns`, `maxOpenConns`, `maxLifetime`,
	// `maxIdleTime`) are empty, it means that the necessary configuration information is missing. In this
	// case, the code logs a message indicating that it is unable to load the database pool configuration.
	if maxIdleConns == "" || maxOpenConns == "" || maxLifetime == "" || maxIdleTime == "" {
		log.Println("Unable to load database pool configuration.")
	}

	// Sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(utils.ToInt(maxOpenConns) | 100)

	// Set the maximum number of idle (unused) connections in the connection pool
	// for the SQL database connection represented.
	db.SetMaxIdleConns(utils.ToInt(maxIdleConns) | 25)

	// Set the maximum amount of time a connection can be reused in the
	// connection pool before it is closed and discarded.
	db.SetConnMaxLifetime(time.Minute*time.Duration(utils.ToInt(maxLifetime)) | 15)

	// Sets the maximum amount of time a connection can remain idle in the
	// connection pool before it is closed and discarded.
	db.SetConnMaxIdleTime(time.Minute*time.Duration(utils.ToInt(maxIdleTime)) | 10)
}
