package main

import (
	"github.com/rizalreza/golang-echo-rest/database"
	"github.com/rizalreza/golang-echo-rest/routes"
)

func main() {

	database.InitialConnection()
	e := routes.InitialRoute()

	e.Logger.Fatal(e.Start(":5000"))
}
