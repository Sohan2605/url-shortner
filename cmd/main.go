package main

import (
	"log"
	"net/http"

	"url-shortener/internal/handler"
	"url-shortener/pkg/db"
	"url-shortener/pkg/kafka"
	"url-shortener/pkg/redis"
)

func main() {
	db.Init()
	redis.Init()
	kafka.Init()

	http.HandleFunc("/shorten", handler.Shorten)
	http.HandleFunc("/", handler.Redirect)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
