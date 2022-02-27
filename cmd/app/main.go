package main

import (
	"log"

	"github.com/gopherzz/gopbin/internal/app"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config := &app.AppConfig{
		DocumentsFolderPath: "docs",
		Port:                ":8080",
	}
	a := app.NewApp(config)
	return a.Start()
}
