package main

import (
	"fmt"
	"sync"
	"time"
)

var str string
var mutex = &sync.Mutex{}
var wg sync.WaitGroup

func routine() {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		str = "inside goroutine"
		fmt.Println(str)
		mutex.Unlock()
	}
}

func doSomething() {
	defer wg.Done()
	fmt.Println("I'm inside the wg.Done func")
}

func main() {
	go routine()
	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			str = "inside main"
			fmt.Println(str)
			mutex.Unlock()
		}
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doSomething()
	}
	time.Sleep(time.Second * 3)
}
