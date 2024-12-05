package main

import (
	"fmt"
	"sync"
)

const workersCount = 3

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	animals := []string{"dog", "cat", "bird", "seal", "cow", "platypus"}
	// Writer
	go func() {
		for _, animal := range animals {
			ch <- animal
		}
		close(ch)
	}()

	wg.Add(workersCount)

	counter := 1
	var mu sync.Mutex

	for w := 1; w <= workersCount; w++ {
		// Reader
		go func(w int) {
			defer wg.Done()
			for animal := range ch {
				mu.Lock()
				// line 1, thread 1, animal "dog"
				fmt.Printf("line %d, thread %d, animal \"%s\"\n", counter, w, animal)
				counter++
				mu.Unlock()
			}
		}(w)
	}

	wg.Wait()
	fmt.Println("done")
}
