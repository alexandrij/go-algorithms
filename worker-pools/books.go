package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sync/singleflight"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Book struct {
	ID    int
	Title string `json:"Title"`
}

func main() {
	s := singleflight.Group{}

	http.HandleFunc("/getBook", func(writer http.ResponseWriter, request *http.Request) {
		bookID, _ := strconv.Atoi(strings.TrimLeft(request.RequestURI, "/getBook"))
		workHash := fmt.Sprintf("book:%d", bookID)

		result, _, _ := s.Do(workHash, func() (interface{}, error) {
			return getBookFromDB(bookID), nil
		})

		book := result.(Book)
		bytes, _ := json.Marshal(book)

		writer.Header().Add("Content-Type", "application/json")
		writer.Write(bytes)

		http.ListenAndServe(":8890", nil)
	})
}

func getBookFromDB(id int) Book {
	time.Sleep(1 * time.Second)
	return Book{ID: id, Title: randSeq(rand.Intn(30))}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
