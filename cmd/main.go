package main

import (
	"log"

	"taskManager/internal"
	"taskManager/internal/db"
)

func main() {
	database, err := db.Open("tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	repo := internal.NewSQLiteRepository(database)

	srv := internal.New()
	srv.RegisterRoutes(repo)

	if err := srv.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
