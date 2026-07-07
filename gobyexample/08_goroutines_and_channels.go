package gobyexample

import (
	"fmt"
	"time"
)

func goroutinesAndChannelsGetValues(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func goroutinesAndChannelsGetValuesWithChannel(ch chan string) {
	for i := range 3 {
		ch <- fmt.Sprintf("Value from channel: %d", i)
	}
	close(ch)
}

func goroutinesAndChannelsWorker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ShowGoroutinesAndChannels() {
	// synchronous call
	goroutinesAndChannelsGetValues("direct")

	// run function in a new goroutine
	go goroutinesAndChannelsGetValues("goroutine")

	// wait for the goroutine to finish
	time.Sleep(time.Second / 2)

	// create a channel to communicate between goroutines
	ch := make(chan string)

	// run function in a new goroutine and pass the channel
	go goroutinesAndChannelsGetValuesWithChannel(ch)

	// read values from the channel
	for val := range ch {
		fmt.Println(val)
	}

	// create a channel to signal when the worker is done
	done := make(chan bool)

	// run the worker in a new goroutine and pass the done channel
	go goroutinesAndChannelsWorker(done)

	// wait for the worker to finish
	<-done
}
