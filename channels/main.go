package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Write a program in which one go routine writes sth to one channel
// Another go routine consumes it
func main() {
	rand.Seed(time.Now().UnixNano())
	sharedVariable := make(chan int, 1)

	var Mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			setRandom(i, &Mutex, sharedVariable)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	wg.Add(2)
	go WriteToChannel(sharedVariable, &wg)
	go ReadFromChannel(sharedVariable, &wg)

	wg.Wait()
}

func setRandom(i int, mut *sync.Mutex, sharedChannel chan int) {
	fmt.Println("Thread #", i)

	// Generate the random number outside the critical section to avoid holding the mutex longer than necessary
	randNum := rand.Int() % 50

	mut.Lock() // Lock before accessing shared resources
	sharedChannel <- randNum
	fmt.Printf("Sent by Thread #%d: %v\n", i, randNum)
	mut.Unlock() // Unlock after accessing shared resources

	// Receive from the channel outside the mutex lock
	receivedNum := <-sharedChannel
	fmt.Printf("Received by Thread #%d: %v\n", i, receivedNum)
}

func WriteToChannel(ch chan<- int, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	ch <- 1
	defer wg.Done()
}

func ReadFromChannel(ch <-chan int, wg *sync.WaitGroup) {
	channeldata := <-ch
	fmt.Println("Read from channel = ", channeldata)
	defer wg.Done()
}
