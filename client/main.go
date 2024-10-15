package main

import (
	"dummy-endpoints/requester"
	"dummy-endpoints/structs"
	"fmt"
	"log"
	"sort"
	"sync"
)

func main() {
	beginPort := structs.GetPorts().Min
	endPort := structs.GetPorts().Max

	var wg sync.WaitGroup
	var mu sync.Mutex
	allResponses := []structs.Response{}

	// Iterate through each port and send requests concurrently
	for port := beginPort; port <= endPort; port++ {
		wg.Add(1)

		go func(port int) {
			defer wg.Done()

			url := fmt.Sprintf("http://localhost:%d", port)
			responses, err := requester.MakeWG(url)
			if err != nil {
				log.Printf("Error for port %d: %v", port, err)
				return
			}

			// Lock before modifying the shared slice
			mu.Lock()
			allResponses = append(allResponses, responses...)
			mu.Unlock()

		}(port)
	}

	// Wait for all Go routines to complete
	wg.Wait()

	// Sort the responses by the Address field
	sort.Slice(allResponses, func(i, j int) bool {
		return allResponses[i].Address < allResponses[j].Address
	})

	// Print or process the sorted responses
	for _, response := range allResponses {
		fmt.Printf("Port Response: %+v\n", response)
	}
}
