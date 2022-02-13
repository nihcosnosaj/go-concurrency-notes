package main

import (
	"fmt"
	"sync"
	"time"
)

func greedy() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	// greedyWorker holds onto the shared lock for the entirety of its work loop.
	// greedyWorker should get almost twice the amount of loops done as politeWorker
	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	// politeWorker only locks when it needs to.
	// politeWorker should only be able to do about half the work loops of greedyWorker
	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()

	wg.Wait()
}

func main() {
	greedy()
}
