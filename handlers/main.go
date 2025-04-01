package handlers

import (
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "content/html/index.html")
}
