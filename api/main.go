package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Connect to Redis service (hostname = "redis" because of docker-compose)
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		val, err := rdb.Incr(ctx, "hits").Result()
		if err != nil {
			http.Error(w, "Redis error: "+err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "Hello from Go! Page visited %d times.\n", val)
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
