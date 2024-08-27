package main

import (
	"fmt"
	"time"
)

// Fire go routines like crazy

func main() {
	f := func() {
		fmt.Println("what the hell")
	}
	go f()
	go DoSthElse()
	go DoSth()
	time.Sleep(time.Second)
}

func DoSth() {
	fmt.Println("I'm lazy, go do it yourself")
}

func DoSthElse() {
	fmt.Println("Don't test my patience!")
}
