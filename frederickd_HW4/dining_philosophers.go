// I modeled my program off of Dijkstra's solution on the Wikipedia page

package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
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

var state [philosophers]State // Keep track of eveyrone's available states
var crit_region_mtx sync.Mutex // For critial regions
var output_mtx sync.Mutex // For printing to terminal
var sem [philosophers]chan struct{}

func rand_int(min, max int) int {
	return rand.Intn(max - min + 1) + min
}

func semaphores(i int) {
	crit_region_mtx.Lock()
	defer crit_region_mtx.Unlock()

	if state[i] == HUNGRY &&
	   state[left(i)] != EATING &&
		 state[right(i)] != EATING {
		select {
		case sem[i] <- struct{}{}:
			state[i] = EATING
		default:
			// Do nothing
		}
	}
}

func think(i int) {
	duration := rand_int(400, 800)
	output_mtx.Lock()
	fmt.Printf("%d is thinking for %dms\n", i, duration)
	time.Sleep(time.Duration(duration) * time.Millisecond)
	output_mtx.Unlock()
}

func take_forks(i int) {
	crit_region_mtx.Lock()
	state[i] = HUNGRY
	semaphores(i)
	crit_region_mtx.Unlock()
	<-sem[i]
}

func eat(i int) {
	duration := rand_int(400, 800)
	output_mtx.Lock()
	fmt.Printf("%d is eating %d", i, duration)
	time.Sleep(time.Duration(duration) * time.Millisecond)
	output_mtx.Unlock()
}

func put_forks(i int) {
	crit_region_mtx.Lock()
	state[i] = THINKING
	semaphores(left(i))
	semaphores(right(i))
	crit_region_mtx.Unlock()
}

func philosopher(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal when a philosopher is finished eating
	for {
		think(i)
		take_forks(i)
		eat(i)
		put_forks(i)
	}
}

func main () {
	var wg sync.WaitGroup

	for i := 0; i < philosophers; i++ {
		sem[i] = make(chan struct{}, 1)
	}

	for i := 0; i < philosophers; i++ {
		wg.Add(1)
		go philosopher(i, &wg)
	}

	wg.Wait()
}

