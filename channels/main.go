package main

import (
	"log"
	"time"
)

func main()  {
	words := make(chan string)

	go say(words)
	log.Println("[MAIN] Hello")

	message := <- words
	log.Println("[MAIN]", message)

	words <- "message from main"

	message2 := <- words
	log.Println("[MAIN]", message2)
}

func say(c chan string) {
	c <- "message 1"
	time.Sleep(time.Second)
	log.Println("[GOROUTINE] World")

	message := <- c
	log.Println("[GOROUTINE]", message)

	c <- "message 2"
}
