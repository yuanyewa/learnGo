package main

import (
	"fmt"
	"sync"
	"time"
)

var aChan = make(chan bool)

func A() {
	fmt.Println("A starts")
	time.Sleep(time.Second * 1)
	fmt.Println("A ends")
	aChan <- true
}

var bChan = make(chan bool)

func B() {
	fmt.Println("B starts")
	time.Sleep(time.Second * 2)
	fmt.Println("B ends")
	bChan <- true
}

var cChan = make(chan bool)

func C() {
	fmt.Println("C starts")
	time.Sleep(time.Second * 1)
	fmt.Println("C ends")
	cChan <- true
}

var dChan = make(chan bool)

func D() {
	<-cChan
	fmt.Println("D starts")
	time.Sleep(time.Second * 1)
	fmt.Println("D ends")
	dChan <- true
}

var eChan = make(chan bool)

func E() {
	fmt.Println("E starts")
	time.Sleep(time.Second * 3)
	eChan <- true
}

var fChan = make(chan bool)

func F() {
	fmt.Println("F starts")
	time.Sleep(time.Second * 1)
	fChan <- true
}

var gChan = make(chan bool)

func G() {
	fmt.Println("G starts")
	time.Sleep(time.Second * 1)
	gChan <- true
}

var hChan = make(chan bool, 1)

func H() {
	fmt.Println("H starts")
	time.Sleep(time.Second * 1)
	hChan <- true
}

var wg sync.WaitGroup

func I() {
	fmt.Println("I starts")
	time.Sleep(time.Second * 1)
	fmt.Println("I ends")
	wg.Done()
}

func J() {
	fmt.Println("I starts")
	time.Sleep(time.Second * 1)
	fmt.Println("I ends")
	wg.Done()
}

var ans int = 0

var mu = &sync.Mutex{}

func K() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		ans++
		time.Sleep(time.Millisecond * 1)
		mu.Unlock()
	}
	wg.Done()
}
func L() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		ans++
		time.Sleep(time.Millisecond * 1)
		mu.Unlock()
	}
	wg.Done()
}

func M() {
	for i := 0; i < 10; i++ {
		ans++
	}
}

func main() {
	// now := time.Now()
	// go A()
	// go B()
	// <-aChan
	// <-bChan
	// fmt.Println(time.Since(now))

	// now := time.Now()
	// go C()
	// go D()
	// <-dChan
	// fmt.Println(time.Since(now))

	// go E()
	// select {
	// case <-eChan:
	// 	fmt.Println("E ends")
	// case <-time.After(time.Second * 2):
	// 	fmt.Println("time out after 2 sec")
	// }

	// go F()
	// go G()
	// select {
	// case <-fChan:
	// 	fmt.Println("F ends")
	// case <-gChan:
	// 	fmt.Println("G ends")
	// }

	// H()
	// <-hChan
	// fmt.Println("H ends")

	// queue := make(chan int, 20)
	// queue <- 1
	// queue <- 2
	// close(queue)
	// for v := range queue {
	// 	fmt.Println(v)
	// }

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()
	fmt.Println("timer starts")
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("timer expires")

	// go I()
	// go J()
	// wg.Add(2)
	// wg.Wait()

	// go K()
	// go L()
	// wg.Add(2)
	// wg.Wait()
	// fmt.Println(ans)
	M()
}
