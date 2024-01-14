package main

import "github.com/muhsenmaqsudi/goravel"

type application struct {
	App *goravel.Goravel
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
