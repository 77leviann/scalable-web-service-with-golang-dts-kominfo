package main

import (
	"log"
	"net/http"
	"time"

	"assignment-3/internal/handler"
	"assignment-3/internal/service"
)

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)

	http.HandleFunc("/status", handler.StatusHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	go func() {
		log.Println("Server started at :8888")
		log.Fatal(http.ListenAndServe(":8888", nil))
	}()

	go func() {
		for {
			log.Println("Updating status...")
			time.Sleep(15 * time.Second)
		}
	}()

	select {}
}
