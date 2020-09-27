package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rizalreza/golang-echo-rest/config"
)

var db *sql.DB
var err error

func InitialConnection() {
	conf := config.GetConfig()

	connection := conf.Username + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.Name

	db, err = sql.Open("mysql", connection)

	if err != nil {
		panic("Connection failed")
	}

	err = db.Ping()

	if err != nil {
		panic("Database not found or invalid credential")
	}
}

func CreateConnection() *sql.DB {
	return db
}
