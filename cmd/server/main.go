package main

import (
	"log"
	"net/http"

	handler "github.com/kevintamakuwala/retail-pulse/internal/api"
)

func main() {
	jobHandler := handler.NewJobHandler()

	http.HandleFunc("/api/submit/", jobHandler.SubmitHandler)
	http.HandleFunc("/api/status/", jobHandler.StatusHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
