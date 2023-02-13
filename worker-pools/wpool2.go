package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type WorkJobResult struct {
	Status bool
}

type WorkJob struct {
	ID     int
	Result chan WorkJobResult
}

type WorkerPool struct {
	jobs chan WorkJob
}

func (w *WorkerPool) StartWorker() {
	go func() {
		for {
			work := <-w.jobs

			time.Sleep(1 * time.Second)

			status := false
			if work.ID%10 > 5 {
				status = true
			}

			work.Result <- WorkJobResult{Status: status}
		}
	}
}

func (w *WorkerPool) AddJob(ctx context.Context, id int) <-chan WorkJobResult {
	resultChan := make(chan WorkJobResult, 1)

	select {
	// trying to add work job
	case w.jobs <- WorkJob{ID: id, Result: resultChan}:
	// resurn chan where consumer can read result
	case <-ctx.Done():
		return nil
	}
	// return chan where consumer can read result
	return resultChan
}

func main() {
	wp := WorkerPool{
		jobs: make(chan WorkJob, 3),
	}
	wp.StartWorker()
	wp.StartWorker()
	wp.StartWorker()

	http.HandleFunc("/handle", func(writer http.ResponseWriter, request *http.Request) {
		resultsChan := make([]<-chan WorkJobResult, 0)
		for i := 0; i < 10; i++ {
			resultChan := wp.AddJob(context.Background(), rand.Intn(100))
			resultsChan = append(resultsChan, resultChan)
		}

		status := false
		for _, res := range resultsChan {
			resStatus := <-res
			status = status && resStatus.Status
		}

		bytes, _ := json.Marshal(status)
		writer.Write(bytes)
	})

	_ = http.ListenAndServe(":8890", nil)
}
