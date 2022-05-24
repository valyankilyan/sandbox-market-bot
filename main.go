package main

import (
	"fmt"
	"playground/helloworld"
	"sync"
	"sync/atomic"
	"time"
)

var m sync.RWMutex

func SayHiMain(hi string) {
	fmt.Println(hi)
}

func Anschanger(ans *string, n, k int, id string) {
	time.Sleep(time.Second)
	for i := 0; i < n; i += k {
		m.Lock()
		*ans = fmt.Sprintf("%v {%d} - %d - %d", id, i, n, k)
		m.Unlock()
		time.Sleep(time.Millisecond * time.Duration(n))
	}
}

func ansreader(ans *string, n int) {
	for i := 0; i < n; i++ {
		m.RLock()
		fmt.Printf("start reader %d\n", n)
		time.Sleep(time.Millisecond * 100 * time.Duration(n))
		fmt.Printf("have read %v\n", *ans)
		time.Sleep(time.Second)
		m.RUnlock()
		fmt.Println("RUNLOCKED", n)
	}
}

func Hello(name string, wait int) {
	time.Sleep(time.Millisecond * 100 * time.Duration(wait))
	fmt.Printf("Hello %v\n", name)
}

func main() {
	ans := "ans"

	// go Anschanger(&ans, 2, 1, "first")
	time.Sleep(time.Millisecond * 100)
	go Anschanger(&ans, 1, 1, "second")
	go ansreader(&ans, 1)
	go ansreader(&ans, 1)

	prevans := ans
	for i := 0; i < 1; i++ {
		m.RLock()
		if prevans != ans {
			fmt.Println(ans)
			prevans = ans
		}
		time.Sleep(time.Millisecond * 500)
		m.RUnlock()
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		Hello("Valya", 5)
		wg.Done()
	}()
	go func() {
		// wg.Wait()
		Hello("B", 3)
		wg.Done()
	}()
	go func() {
		// wg.Wait()
		Hello("C", 1)
		wg.Done()
	}()
	fmt.Println("WAITING")
	wg.Wait()
	fmt.Println("DONE WAITING")
	// time.Sleep(time.Second * 100)

	var ac, bc, wc int64
	ac, bc, wc = 0, 0, 0
	// var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				ac++
				bc++
				// wc++
				atomic.AddInt64(&wc, 1)
				// time.Sleep(time.Microsecond)
				// wc++
			}
			wg.Done()
		}()
	}
	for i := 0; i < 100000; i++ {
		atomic.AddInt64(&wc, 1)
	}
	wg.Wait()
	fmt.Println(ac, bc, wc)

	helloworld.SayHi("Name")
}
