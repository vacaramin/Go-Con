package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sharedVariable int = 2

func main() {
	rand.Seed(time.Now().UnixNano())
	var Mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			setRandom(i%50, &Mutex)
			wg.Done()
		}()
	}
	wg.Wait()
}

func setRandom(i int, mut *sync.Mutex) {
	mut.Lock()
	sharedVariable = rand.Int()
	time.Sleep(time.Second)
	mut.Unlock()
	fmt.Printf("%v\n", i)
}
