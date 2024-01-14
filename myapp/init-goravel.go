package main

import (
	"log"
	"os"

	"github.com/muhsenmaqsudi/goravel"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cel := &goravel.Goravel{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myApp"

	cel.InfoLog.Println("Debug is set to", cel.Debug)

	app := &application{
		App: cel,
	}

	return app
}
