package main

import (
	"myapp/data"
	"myapp/handlers"

	"github.com/muhsenmaqsudi/goravel"
)

type application struct {
	App      *goravel.Goravel
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
