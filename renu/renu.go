package main

//Buffer channels

// // 1 // consumer is ready
// func main() {
// 	channel := make(chan string)
// 	go func() {
// 		channel <- "First message"
// 	}()
// 	fmt.Println(<-channel)
// }

// // 2 // since channel size is 0
// func main() {
// 	channel := make(chan string)
// 	channel <- "First message"
// 	fmt.Println(<-channel)
// }

// // 3
// func main() {
// 	channel := make(chan string, 1)
// 	channel <- "First message"
// 	fmt.Println(<-channel)
// }

// // 4 using Go buffer
// func main() {
// 	channel := make(chan string, 1)

// 	go func() {
// 		fmt.Println("In go routine", <-channel)
// 	}()

// 	channel <- "First message"
// 	fmt.Println("12 ")

// 	channel <- "Second message"
// 	fmt.Println("22")
// 	fmt.Println(<-channel)

// }

// // 5
// func main() {
// 	channel := make(chan string, 2)

// 	channel <- "First message"
// 	channel <- "Second message"

// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)
// }

//channels gives FIFO, what if I want to iterate over it
