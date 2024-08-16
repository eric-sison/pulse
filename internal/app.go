package internal

import (
	"database/sql"
	"log"

	"github.com/eric-sison/pulse/pkg/utils"
)

type App struct {
	Db *sql.DB
}

func CreateApp() App {
	utils.LoadEnv()

	db, err := ConnectDB()

	if err != nil {
		log.Fatal()
	}

	app := App{Db: db}

	return app
}
