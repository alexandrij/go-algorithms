package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Movie struct {
	Title string `json:"Title"`
}

type CachedPopularItems struct {
	lock   sync.RWMutex
	Movies []Movie
}

func main() {
	ctx := context.Background()

	cache := CachedPopularItems{}
	cache.Movies = getPopularMoviesFromDB()

	go func() {
		timer := time.NewTicker(1 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				movies := getPopularMoviesFromDB()

				cache.lock.Lock()
				cache.Movies = movies
				cache.lock.Unlock()

			case <-ctx.Done():
				break
			}
		}
	}()

	http.HandleFunc("/getPopularMovies", func(writer http.ResponseWriter, request *http.Request) {
		cache.lock.RLock()
		movies := cache.Movies
		cache.lock.RUnlock()

		bytes, _ := json.Marshal(movies)

		writer.Header().Add("Content-Type", "application/json")
		writer.Write(bytes)
	})

	_ = http.ListenAndServe(":8890", nil)
}

func getPopularMoviesFromDB() []Movie {
	time.Sleep(5 * time.Second)

	return []Movie{
		{Title: "Avatar"},
		{Title: "I am Legend"},
		{Title: "The Wolf of Wall Street"},
	}
}
