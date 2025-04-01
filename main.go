package main

import (
	"fmt"
	"log"
	"net/http"
	"smppmainhttpserver/handlers"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/roadmap", handlers.RoadmapHandler)
	port := "9000"
	fmt.Println("Starting HTTP server on port: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
