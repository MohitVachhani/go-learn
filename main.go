package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func squareIt(wg *sync.WaitGroup, numbersChannel chan int, workerID int, doneChannels *map[int]int, mutex *sync.Mutex, ctx context.Context) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case num := <-numbersChannel:
// 			{
// 				result := num * num
// 				mutex.Lock()
// 				(*doneChannels)[workerID]++
// 				mutex.Unlock()
// 				fmt.Printf("Worker %d processed %d -> %d\n", workerID, num, result)
// 			}
// 		case <-ctx.Done():
// 			{
// 				return
// 			}
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	numbersChannel := make(chan int)
// 	doneChannels := make(map[int]int)
// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go squareIt(&wg, numbersChannel, i, &doneChannels, &mutex, ctx)
// 	}

// 	// Send numbers to workers
// 	go func() {
// 		defer close(numbersChannel)
// 		for i := 0; i < 1000; i++ {
// 			select {
// 			case <-ctx.Done():
// 				// Stop sending if context timeout happens
// 				return
// 			case numbersChannel <- i:
// 			}
// 		}
// 	}()

// 	wg.Wait()
// 	for key, iterator := range doneChannels {
// 		fmt.Printf("Worker %d processed %d numbers \n", key, iterator)
// 	}
// }

import (
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	conf := pdfcpu.NewDefaultConfiguration()

	err := api.ExtractImagesFile("page_1_modern.pdf", "./out", nil, )
	if err != nil {
		log.Fatal("No images found or extraction failed:", err)
	}
}
