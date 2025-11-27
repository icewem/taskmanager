package main

import (
	"taskManager/internal"
)

func main() {
	srv := internal.New()
	srv.RegisterRoutes()
	srv.Start(":8080")
}
