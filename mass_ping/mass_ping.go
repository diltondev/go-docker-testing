package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

const NUM_ITER = 1000

func main() {
	fmt.Println("Sending mass pings...")
	MassPing()
}

func MassPing() {
	wg := sync.WaitGroup{}

	wg.Add(NUM_ITER)
	for i := 0; i < NUM_ITER; i++ {
		go func(in int) {
			defer wg.Done()
			fmt.Printf("Sent ping %d\n", in)
			res, err := http.Get("http://localhost:8080/ping")
			if err != nil {
				fmt.Println("Error when performing GET request")
				return
			}
			body, _ := io.ReadAll(res.Body)
			fmt.Printf("Body for request %d:\n    %s", in, body)
		}(i)
	}
	wg.Wait()
	fmt.Println("finished requests")
}
