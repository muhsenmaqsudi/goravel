package handlers

import (
	"fmt"
	"myapp/data"
	"net/http"

	"github.com/muhsenmaqsudi/goravel"
)

type Handlers struct {
	App *goravel.Goravel
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Welcome to myApp")
	if err != nil {
		h.App.ErrorLog.Println("error:", err)
	}
}
