package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var count uint64

func main() {
	num := os.Getenv("SERVER_NUMBER")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&count, 1)
		fmt.Fprintf(w, "Hello from server number %v!", num)
	})

	go func() {
		t := time.NewTicker(10 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				current := atomic.LoadUint64(&count)
				fmt.Printf("amount of requests to server number %v is %v\n", num, current)
			}
		}
	}()

	http.ListenAndServe(":8088", nil)
}
