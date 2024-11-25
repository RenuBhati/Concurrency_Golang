package main

// //3 these create deadlock and deadlock can be solved using buffer
// func main() {
// 	now := time.Now()
// 	defer func() {
// 		fmt.Println(time.Since(now))
// 	}()
// 	smokeSignal := make(chan bool)
// 	evilNinja := "Tommy"
// 	go attack(evilNinja, smokeSignal)
// 	smokeSignal <- false
// 	fmt.Println(<-smokeSignal)
// }

// func attack(target string, attacked chan bool) {
// 	time.Sleep(time.Second * 1)
// 	fmt.Println("Throwing ninja starts at", target)
// 	attacked <- true
// }

// //2
// func main() {
// 	now := time.Now()
// 	defer func() {
// 		fmt.Println(time.Since(now))
// 	}()

// 	smokeSignal := make(chan bool)
// 	evilNinja := "Tommy"
// 	go attack(evilNinja, smokeSignal)
// 	fmt.Println(<-smokeSignal)
// }

// func attack(target string, attacked chan bool) {
// 	time.Sleep(time.Second * 1)
// 	fmt.Println("Throwing ninja starts at", target)
// 	attacked <- true
// }

// //1
// func main() {
// 	now := time.Now()
// 	defer func() {
// 		fmt.Println(time.Since(now))
// 	}()
// 	evilNinja := "Tommy"
// 	go attack(evilNinja)
// 	time.Sleep(time.Second * 3)
// }

// func attack(target string) {
// 	time.Sleep(time.Second * 1)
// 	fmt.Println("Throwing ninja starts at", target)
// }
