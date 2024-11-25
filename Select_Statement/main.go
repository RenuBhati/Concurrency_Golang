package main

// The select statement lets a goroutine wait on multiple communication operations
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready

// //1
// func main() {
// 	ninja1, ninja2 := make(chan string), make(chan string)

// 	go captainElect(ninja1, "Hello ninja11")
// 	go captainElect(ninja2, "Two Hello ninja11")

// 	fmt.Println(<-ninja1)
// 	fmt.Println(<-ninja2)
// }

// func captainElect(ninja chan string, message string) {
// 	ninja <- message
// }

// // 2
// func main() {
// 	ninja1, ninja2 := make(chan string), make(chan string)

// 	go captainElect(ninja1, "Hello ninja11")
// 	go captainElect(ninja2, "Two Hello ninja11")

// 	select {
// 	case message := <-ninja1:
// 		fmt.Println(message)
// 	case message := <-ninja2:
// 		fmt.Println(message)

// 	}
// }

// func captainElect(ninja chan string, message string) {
// 	ninja <- message
// }

// 3 Rughly fair

// func main() {
// 	ninja1, ninja2 := make(chan string), make(chan string)

// 	go captainElect(ninja1, "Hello ninja11")
// 	go captainElect(ninja2, "Two Hello ninja11")

// 	select {
// 	case message := <-ninja1:
// 		fmt.Println(message)
// 	case message := <-ninja2:
// 		fmt.Println(message)
// 	default:
// 		fmt.Println("Neither")
// 	}
// 	roughlyFair()
// }

// func captainElect(ninja chan string, message string) {
// 	ninja <- message
// }

// func roughlyFair() {
// 	ninja1 := make(chan interface{})
// 	close(ninja1)
// 	ninja2 := make(chan interface{})
// 	close(ninja2)

// 	var ninja1Count, ninja2Count int
// 	for i := 0; i < 1000; i++ {
// 		select {
// 		case <-ninja1:
// 			ninja1Count++
// 		case <-ninja2:
// 			ninja2Count++

// 		}
// 	}
// 	fmt.Printf("ninja1Count: %d, ninja2Count: %d\n", ninja1Count, ninja2Count)
// }
