package main

// 1

// func main() {
// 	go attack("Tommy")
// 	fmt.Println("MISSION COMPLETE")
// }

// func attack(evilNinja string) {
// 	fmt.Println("Attacked evil ninja:", evilNinja)
// }

// // 2
// func main() {
// 	var beeper sync.WaitGroup
// 	beeper.Add(1)
// 	go attack("Tommy", &beeper)
// 	beeper.Wait()
// 	fmt.Println("MISSION COMPLETE")
// }

// func attack(evilNinja string, beeper *sync.WaitGroup) {
// 	fmt.Println("Attacked evil ninja:", evilNinja)
// 	beeper.Done()
// }

// // 3
// func main() {
// 	var beeper sync.WaitGroup
// 	evilNinjas := [] string{"Tommy", "Johnny", "Bobby"}
// 	beeper.Add(len(evilNinjas))
// 	for _, evilNinja := range evilNinjas {
// 	go attack(evilNinja, &beeper)
// 	}
// 	beeper.Wait()
// 	fmt.Println("MISSION COMPLETE")
// }

// func attack(evilNinja string, beeper *sync.WaitGroup) {
// 	fmt.Println("Attacked evil ninja:", evilNinja)
// 	beeper.Done()
// }

// // 4
// func main() {
// 	var beeper sync.WaitGroup
// 	beeper.Add(1)
// 	beeper.Wait()
// }

// // 5
// func main() {
// 	var beeper sync.WaitGroup
// 	beeper.Add(1)
// 	beeper.Done()
// 	beeper.Wait()
// }

// // 5
// func main() {
// 	var beeper sync.WaitGroup
// 	beeper.Add(1)
// 	beeper.Done()
// 	beeper.Done()
// 	beeper.Wait()
// }
