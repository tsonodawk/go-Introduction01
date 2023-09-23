package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sayHallo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hallo")
}

func memConsumed() uint64 {
	runtime.GC()

	var s runtime.MemStats

	runtime.ReadMemStats(&s)

	return s.Sys
}

func Hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %v\n", id)
}

type value struct {
	mu    sync.Mutex
	value int
	name  string
}

func main() {
	// Basic()
	// closure()
	// memory()
	// waitGroup()
	// syncMutex()
	syncDeadlock()
}

func syncDeadlock() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		fmt.Printf("%v がロックしました\n", v1.name)
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)

		v2.mu.Lock()
		fmt.Printf("%v がロックしました\n", v2.name)
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a value = value{name: "a"}
	var b value = value{name: "b"}

	wg.Add(2)

	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()

}

func syncMutex() {
	var wg sync.WaitGroup
	var memoryAccess sync.Mutex
	var data int

	wg.Add(1)
	go func() {
		defer wg.Done()
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	wg.Wait()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}

func waitGroup() {

	var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("1st goroutine sleeping...")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("1st goroutine Done")

	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("2st goroutine sleeping...")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("2st goroutine Done")

	// }()

	// wg.Wait()

	var CPU int = runtime.NumCPU()

	wg.Add(CPU)
	for i := 0; i < CPU; i++ {
		go Hello(&wg, i)
	}

	wg.Wait()

}

func memory() {

	var ch <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-ch
	}

	const numGoroutines = 1000000

	wg.Add(numGoroutines)

	before := memConsumed()

	for i := 0; i < numGoroutines; i++ {
		go noop()
	}

	wg.Wait()

	after := memConsumed()

	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)

}

func closure() {
	var wg sync.WaitGroup

	// say := "Hello"
	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	say = "Good Bye"
	// }()

	// wg.Wait()

	// fmt.Println(say)

	tasks := []string{"task1", "task2", "task3"}
	for _, task := range tasks {
		wg.Add(1)

		go func(task string) {
			defer wg.Done()
			fmt.Println(task)
		}(task)
	}

	wg.Wait()

}

func Basic() {
	var wg sync.WaitGroup

	wg.Add(1)

	go sayHallo(&wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hallo02")
	}()

	// time.Sleep(2 * time.Second)
	wg.Wait()

}
