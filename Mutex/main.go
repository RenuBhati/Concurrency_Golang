package main

import (
	"fmt"
	"sync"
	"time"
)

// // 1
// var (
// 	count int
// )

// func main() {
// 	iterations := 100
// 	for i := 0; i < iterations; i++ {
// 		go increment()
// 	}
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Resulted count is:", count)
// }

// func increment() {
// 	count++
// }

// // 2
// var (
// 	lock  sync.Mutex
// 	count int
// )

// func main() {
// 	iterations := 1000
// 	for i := 0; i < iterations; i++ {
// 		go increment()
// 	}
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Resulted count is:", count)
// }

// func increment() {
// 	lock.Lock()
// 	count++
// 	lock.Unlock()
// }

// 3
var (
	lock   sync.Mutex
	rwLock sync.RWMutex
	count  int
)

func main() {
	// basics()
	readAndWrite()
}

func basics() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is:", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}

func readAndWrite() {
	go read()
	go write()
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()
	fmt.Println("READ LOCKING")
	time.Sleep(1 * time.Second)
	fmt.Println("Reading unlocking")
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()
	fmt.Println("Write LOCKING")
	time.Sleep(1 * time.Second)
	fmt.Println("Write unlocking")
}
