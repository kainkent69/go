package main

import (
	"fmt"
	"log"
)

func Example_channels_recieve_any() {
	log.Println("See if channels cause a constant call of a function over and over")
	c := make(chan string)
	for i := range 10 {
		go func() {
			content := fmt.Sprintf("Example channel: recieved [%d]\n", i)
			c <- content
		}()
	}
	for range 10 {
		select {
		case content := <-c:
			fmt.Print(content)
			break

		}
	}
}
