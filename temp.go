package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Get the default number of threads Go will use
	defaultProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("Default GOMAXPROCS: %d\n", defaultProcs)
}
