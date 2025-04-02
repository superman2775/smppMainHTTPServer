package main

import (
	"fmt"
	"log"
	"net/http"
	"smppmainhttpserver/handler"
)

func main() {
	http.HandleFunc("/", handler.DynamicHandler)
	http.HandleFunc("/set-mode", handler.SetModeHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./content/css"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./content/static"))))
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("./content/media"))))

	port := "9000"
	fmt.Println("Starting Smartschool++ HTTP server on port: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
