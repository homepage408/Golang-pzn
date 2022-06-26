package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// default values from nill return pull
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Teguh")
	pool.Put("Setiawan")
	pool.Put("Budi")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
