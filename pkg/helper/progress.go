package helper

import (
	"fmt"
	"sync"
	"time"
)

func Progressor(ch chan uint8) {

	prog := float32(0)
	mtx := &sync.Mutex{}
	var wg sync.WaitGroup
	totalBlock := 4

	for i := 0; i < totalBlock; i++ {
		wg.Add(1)
		go func(prog *float32, mtx *sync.Mutex) {
			time.Sleep(5 * time.Second)
			defer wg.Done()
			mtx.Lock()
			defer mtx.Unlock()
			fmt.Println("before", *prog)
			*prog += 100.0 / float32(totalBlock)
			fmt.Println("after", *prog)
			ch <- uint8(*prog)
		}(&prog, mtx)
	}

	wg.Wait()
	close(ch)
}
