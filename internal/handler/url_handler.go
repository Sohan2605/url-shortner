package handler

import (
	"fmt"
	"net/http"

	"url-shortener/internal/service"
	"url-shortener/pkg/kafka"
	"url-shortener/pkg/redis"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	code, err := service.CreateShortURL(url)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", code)))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]

	val, err := redis.Client.Get(redis.Ctx, code).Result()
	if err == nil {
		http.Redirect(w, r, val, 302)
		return
	}

	url, err := service.GetOriginalURL(code)
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	redis.Client.Set(redis.Ctx, code, url, 0)
	kafka.Publish(code)

	http.Redirect(w, r, url, 302)
}
