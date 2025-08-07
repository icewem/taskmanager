package main

import (
	"log"
	"taskmanager/internal/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
