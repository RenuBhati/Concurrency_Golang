package main

//Golang Channel Iteration and Channel Closing

// // 1
// func main() {
// 	channel := make(chan string)
// 	go throwingNinjaStar(channel)
// 	fmt.Println(<-channel)
// }

// func throwingNinjaStar(channel chan string) {
// 	// rand.Seed(time.Now().UnixNano())
// 	score := rand.Intn(10)
// 	channel <- fmt.Sprint("You scored ", score)
// }

// // 2
// func main() {
// 	channel := make(chan string)
// 	numRounds := 3
// 	go throwingNinjaStar(channel, numRounds)
// 	for i := 0; i < numRounds; i++ {
// 		fmt.Println(<-channel)
// 	}
// }

// func throwingNinjaStar(channel chan string, numRounds int) {
// 	// rand.Seed(time.Now().UnixNano())

// 	for i := 0; i < numRounds; i++ {
// 		score := rand.Intn(10)
// 		channel <- fmt.Sprint("You scored ", score)
// 	}
// }

// 3 even after 3 loops its doesn't stop and leads to deadlock
// func main() {
// 	channel := make(chan string)

// 	go throwingNinjaStar(channel)
// 	for message := range channel {
// 		fmt.Println(message)
// 	}
// }

// func throwingNinjaStar(channel chan string) {

// 	numRounds := 3
// 	for i := 0; i < numRounds; i++ {
// 		score := rand.Intn(10)
// 		channel <- fmt.Sprint("You scored ", score)
// 	}
// }

// // 4 How to come out of deadlock, even after 3 loop, it will check for 4th and Println is waiting , leading to deadlock
// func main() {
// 	channel := make(chan string)

// 	go throwingNinjaStar(channel)
// 	for message := range channel {
// 		fmt.Println(message)
// 	}
// }

// func throwingNinjaStar(channel chan string) {

// 	numRounds := 3
// 	for i := 0; i < numRounds; i++ {
// 		score := rand.Intn(10)
// 		channel <- fmt.Sprint("You scored ", score)
// 	}
// 	close(channel)
// }

// // 5
// func main() {
// 	channel := make(chan string)

// 	go throwingNinjaStar(channel)
// 	for {
// 		message, open := <-channel
// 		if !open {
// 			break
// 		}
// 		fmt.Println(message)
// 	}
// }

// func throwingNinjaStar(channel chan string) {

// 	numRounds := 3
// 	for i := 0; i < numRounds; i++ {
// 		score := rand.Intn(10)
// 		channel <- fmt.Sprint("You scored ", score)
// 	}
// 	close(channel)
// }
