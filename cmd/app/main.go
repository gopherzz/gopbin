package main

import (
	"log"
	"os"

	"github.com/gopherzz/gopbin/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	config := &app.AppConfig{
		DocumentsFolderPath: os.Getenv("DOCUMENTS_FOLDER_PATH"),
		Port:                ":" + os.Getenv("PORT"),
	}

	a := app.NewApp(config)
	return a.Start()
}
