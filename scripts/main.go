package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	url := "http://localhost:8080/post/event/1?data=first-message"
	succeded := 0
	failed := 0

	var mu sync.Mutex
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Create a new request
			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				fmt.Printf("Request %d: Error creating request: %v\n", i, err)
				return
			}
			req.Header.Set("Content-Type", "application/json")

			// Send the request
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("Request %d: Error sending request: %v\n", i, err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				mu.Lock()
				failed += 1
				mu.Unlock()
			} else {
				mu.Lock()
				succeded += 1
				mu.Unlock()
			}

			fmt.Printf("Request %d: Response status code: %d\n", i, resp.StatusCode)
		}(i)
	}

	// Wait for all requests to finish
	wg.Wait()
	fmt.Printf("All requests completed. succeded: %d, failed: %d\n", succeded, failed)
}
