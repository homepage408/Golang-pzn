package golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// memasukan data ke channel
// channel <- data

// mengambil data dari channel ke variable
// data := <- channel
// fmt.Println(<- channel)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)

		channel <- "Teguh Setiawan"
		fmt.Println("Selesai Mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)

	channel <- "Teguh Setiawan"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// mengirim data ke channel saja
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Teguh Setiawan"
}

// menerima data dari channel saja
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

//buffered channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// channel <- "Teguh"
	// channel <- "setiawan"
	// channel <- "Budi"

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	func() {
		channel <- "Teguh"
		channel <- "setiawan"
		channel <- "Budi"
	}()

	func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	// masuk satu satu selama 10 kali
	go func() {
		for i := 0; i < 10; i++ {
			// fmt.Println(i)
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data ", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		}

		if counter == 2 {
			break
		}

	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}

	}
}
