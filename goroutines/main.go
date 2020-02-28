package main

import (
	"fmt"
	"math/rand"
	"sync"
)
/**
 * Channels can block in two cases:
 * 1. When sending, that will allow two go routines to sync
 * 2. When receiving, when nothing to receive, go routine will be blocked
 *
 * Two types of channel:
 * 1. Buffered - Buffered channels work similar to unbuffered channels, but with one catch — we can send multiple pieces of data to the channel before needing another go routine to read from it.
 * 2. Unbuffered - What makes them unique is that only one piece of data fits through the channel at a time.
 *
 *
 * The main function indeed runs in its own go routine! Even more important to know is that once the main function returns,
 * it closes all other go routines that are currently running.
 * This is why we had a timer at the bottom of our main function — which created a channel and sent a value on it after 5 seconds.
 */
func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChan := make(chan string)
	minedOreChan := make(chan string)

	doneChan := make(chan string)

	//Finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChan <- item
			}
		}
		close(oreChan)
	}(theMine)

	//ore breaker
	go func() {
		for foundOne := range oreChan{
			fmt.Println("Miner: received " + foundOne + " from finder")
			minedOreChan <- foundOne
		}
	}()

	/**
	 * Note: Ranging over a channel will block until another item is sent on the channel.
	 * The only way to stop the go routine from blocking after all sends have occurred is by closing the channel with ‘close(channel)’
	 */

	//Smelter
	go func() {
		for miniedOre := range minedOreChan{
			fmt.Println("From miner: ", miniedOre)
			fmt.Println("From Smelter: ore is smelted")
		}
		doneChan <- "Done!"
	}()

	<- doneChan
}

/*
here is a technique where you can make a non-blocking read on a channel, using Go’s select case structure.
By using the structure below, your go routine will read from the channel if there’s something there, or run the default case
 */
func example02() {

	myChan := make(chan string)

	go func() {
		myChan <- "Message"
	}()

	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No msg")

	}
}
func finder1(data [5]string)  {

	for _, s := range data {
		if s == "ore" {
			fmt.Println("Finder 1 found ore")
		}
	}

}

func finder2(data [5]string)  {
	for _, s := range data {
		if s == "ore" {
			fmt.Println("Finder 2 found ore")
		}
	}
}

func readNumber() int {
	return rand.Intn(10)
}

func example01() {
	var total int
	var wg sync.WaitGroup

	for i := 0; i<10; i++ {
		wg.Add(1)
		go func() {
			for j:= 0; j < 1000; j++ {
				total += readNumber()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}