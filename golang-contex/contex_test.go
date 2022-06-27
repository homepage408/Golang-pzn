package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	backgroun := context.Background()
	fmt.Println(backgroun)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContexWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()

	return destination
}

func TestContexWithCancle(t *testing.T) {
	fmt.Println("Total Go Routine :", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context
	time.Sleep(1 * time.Second)
	fmt.Println("Total Go Routine :", runtime.NumGoroutine())
}

func TestContexWithTimeout(t *testing.T) {
	fmt.Println("Total Go Routine :", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("counter", n)
	}

	// cancel() // mengirim sinyal cancel ke context
	// time.Sleep(1 * time.Second)
	fmt.Println("Total Go Routine :", runtime.NumGoroutine())
}

func TestContexWithDeadline(t *testing.T) {
	fmt.Println("Total Go Routine :", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("counter", n)
	}

	// cancel() // mengirim sinyal cancel ke context
	// time.Sleep(1 * time.Second)
	fmt.Println("Total Go Routine :", runtime.NumGoroutine())
}
