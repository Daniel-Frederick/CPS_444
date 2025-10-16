// I modeled my program off of Dijkstra's solution on the Wikipedia page
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const philosophers = 5

// state enum
type State int

const (
	THINKING State = iota
	HUNGRY
	EATING
)

func left(i int) int {
	return (i + philosophers - 1) % philosophers
}

func right(i int) int {
	return (i + 1) % philosophers
}

var state [philosophers]State       // Keep track of everyone's available states
var crit_region_mtx sync.Mutex      // For critical regions
var output_mtx sync.Mutex           // For printing to terminal
var sem [philosophers]chan struct{} // Semaphores for each philosopher

func rand_int(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// test checks if philosopher i can eat
func test(i int) {
	if state[i] == HUNGRY &&
		state[left(i)] != EATING &&
		state[right(i)] != EATING {
		state[i] = EATING
		// Signal that philosopher i can eat
		select {
		case sem[i] <- struct{}{}:
		default:
			// Channel already has a signal
		}
	}
}

func think(i int) {
	duration := rand_int(400, 800)
	output_mtx.Lock()
	fmt.Printf("%d is thinking for %dms\n", i, duration)
	output_mtx.Unlock()
	time.Sleep(time.Duration(duration) * time.Millisecond)
}

func take_forks(i int) {
	crit_region_mtx.Lock()
	state[i] = HUNGRY
	test(i) // Try to acquire both forks
	crit_region_mtx.Unlock()
	<-sem[i] // Block if forks were not acquired
}

func eat(i int) {
	duration := rand_int(400, 800)
	output_mtx.Lock()
	fmt.Printf("%d is eating for %dms\n", i, duration)
	output_mtx.Unlock()
	time.Sleep(time.Duration(duration) * time.Millisecond)
}

func put_forks(i int) {
	crit_region_mtx.Lock()
	state[i] = THINKING
	// Check if neighbors can now eat
	test(left(i))
	test(right(i))
	crit_region_mtx.Unlock()
}

func philosopher(i int, wg *sync.WaitGroup, iterations int) {
	defer wg.Done()
	for j := 0; j < iterations; j++ {
		think(i)
		take_forks(i)
		eat(i)
		put_forks(i)
	}
	output_mtx.Lock()
	fmt.Printf("Philosopher %d finished\n", i)
	output_mtx.Unlock()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	// Initialize semaphores
	for i := 0; i < philosophers; i++ {
		sem[i] = make(chan struct{}, 1)
	}

	// Start philosophers (each eats 5 times then stops)
	iterations := 5
	for i := 0; i < philosophers; i++ {
		wg.Add(1)
		go philosopher(i, &wg, iterations)
	}

	wg.Wait()
	fmt.Println("All philosophers finished dining")
}
