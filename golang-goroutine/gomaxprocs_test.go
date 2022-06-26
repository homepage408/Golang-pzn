package golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU : ", totalCPU)

	runtime.GOMAXPROCS(10)
	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads : ", totalThreads)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total Go Routine : ", totalGoRoutine)

	group.Wait()
}
