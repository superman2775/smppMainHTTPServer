package handlers

import (
	"net/http"
)

func RoadmapHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "content/html/roadmap.html")
}
