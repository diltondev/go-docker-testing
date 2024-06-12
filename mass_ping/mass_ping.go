package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Sending mass pings...")
	MassPing()
}

func MassPing() {
	for i := 0; i < 1000; i++ {
		go func(in int) {
			fmt.Printf("Sent ping %d\n", in)
			http.Get("http://localhost:8080/ping")
		}(i)
	}
}
