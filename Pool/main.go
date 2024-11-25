package main

// Pool : creating a fixed number or a pool of things for use
// commonly used to constraint things that are expensive

// // 1
// func main() {
// 	var numMemPieces int
// 	memPool := &sync.Pool{
// 		New: func() interface{} {
// 			numMemPieces++
// 			mem := make([]byte, 1024)
// 			return &mem
// 		},
// 	}

// 	const numWorkers = 1024 * 1024

// 	var wg sync.WaitGroup
// 	wg.Add(numWorkers)
// 	for i := 0; i < numWorkers; i++ {
// 		go func() {
// 			mem := memPool.Get().(*[]byte)
// 			memPool.Put(mem)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("%d numMemPieces were created", numMemPieces)
// }

// // 2
// func main() {
// 	var numMemPieces int
// 	memPool := &sync.Pool{
// 		New: func() interface{} {
// 			numMemPieces++
// 			mem := make([]byte, 1024)
// 			return &mem
// 		},
// 	}

// 	const numWorkers = 1024 * 1024

// 	var wg sync.WaitGroup
// 	wg.Add(numWorkers)
// 	for i := 0; i < numWorkers; i++ {
// 		go func() {
// 			memPool.Get()
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("%d numMemPieces were created", numMemPieces)
// }

// // 3
// func main() {
// 	var numMemPieces int
// 	memPool := &sync.Pool{
// 		New: func() interface{} {
// 			numMemPieces++
// 			mem := make([]byte, 1024)
// 			return &mem
// 		},
// 	}

// 	const numWorkers = 1024 * 1024

// 	var wg sync.WaitGroup
// 	wg.Add(numWorkers)
// 	for i := 0; i < numWorkers; i++ {
// 		go func() {
// 			mem := memPool.Get().(*[]byte)
// 			fmt.Sprintln("TAING SOME TIME ON THE RESOURCES")
// 			memPool.Put(mem)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("%d numMemPieces were created", numMemPieces)
// }
