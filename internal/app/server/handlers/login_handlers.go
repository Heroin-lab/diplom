package handlers

import (
	"io"
	"net/http"
)

func HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Vova ebash ih blyat!")
	}
}
