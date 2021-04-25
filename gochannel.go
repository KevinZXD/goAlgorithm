package main

import (
	"fmt"
	"sync"
)

func test13() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	limit := 26

	numChan := make(chan int, 1)
	charChan := make(chan int, 1)
	charChan <- 1

	go func() {
		defer wg.Done()
		for i := 0; i < limit; i++ {
			<-charChan
			fmt.Printf("%c\n", 'a'+i)
			numChan <- 1

		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < limit; i++ {
			<-numChan
			fmt.Println(i)
			charChan <- 1

		}
	}()

	wg.Wait()

	close(charChan)
	close(numChan)
}
func test14() {
	wg := &sync.WaitGroup{}

	for {
		wg.Add(2)
		go func() {
			println(1)
			defer wg.Done()
		}()
		go func() {
			println(2)
			defer wg.Done()
		}()
		wg.Wait()
	}



}
func main() {
	test14()

}
