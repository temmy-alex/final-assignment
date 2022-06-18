package main

import (
	"github.com/temmy-alex/final-assignment/database"
	"github.com/temmy-alex/final-assignment/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
