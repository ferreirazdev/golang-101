package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	urls := []string{"https://go.dev", "https://golang.org", "https://pkg.go.dev"}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			// Simulate HTTP GET
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Fetched: %s\n", u)
		}(url)
	}

	wg.Wait()
	fmt.Println("All fetches done.")
}
