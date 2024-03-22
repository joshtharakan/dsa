package main

import (
	"fmt"
	"sync"
)

func simulatePingPong(rounds int) {

	ping, pong := make(chan bool), make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(2) // Prepare to wait for two goroutines

	go func() {
		defer wg.Done()
		for i := 0; i < rounds; i++ {
			fmt.Println("ping")
			ping <- true
			<-pong
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < rounds; i++ {
			<-ping
			fmt.Println("pong")
			pong <- true
		}
	}()

	wg.Wait()
	close(ping)
	close(pong)
}

func main() {
	simulatePingPong(5)
}
