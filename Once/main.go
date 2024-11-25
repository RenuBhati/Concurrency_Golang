package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Once is an object that will perform exactly one action
// A Once must not copied after first use

// // 1
// var missionCompleted bool

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(100)
// 	for i := 0; i < 100; i++ {
// 		go func() {

// 			if foundTreasure() {
// 				markMissionCompleted()
// 			}
// 			wg.Done()
// 		}()

// 	}
// 	wg.Wait()
// 	checkMissionCompletion()
// }

// func checkMissionCompletion() {
// 	if missionCompleted {
// 		fmt.Println("MISSION IS NOW COMPLETED")
// 	} else {
// 		fmt.Println("MISSION IS A FAILURE")
// 	}
// }

// func markMissionCompleted() {
// 	missionCompleted = true
// }

// func foundTreasure() bool {
// 	rand.Seed(time.Now().UnixNano())
// 	return 0 == rand.Intn(10)
// }

// 2
var missionCompleted bool

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	var once sync.Once
	for i := 0; i < 100; i++ {
		go func() {

			if foundTreasure() {
				once.Do(markMissionCompleted)
			}
			wg.Done()
		}()

	}
	wg.Wait()
	checkMissionCompletion()
}

func checkMissionCompletion() {
	if missionCompleted {
		fmt.Println("MISSION IS NOW COMPLETED")
	} else {
		fmt.Println("MISSION IS A FAILURE")
	}
}

func markMissionCompleted() {
	missionCompleted = true
	fmt.Println("Congrats mission done")
}

func foundTreasure() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
