package main

import (
	"log"
	"time"
)

func main()  {
	channel := make(chan string, 2)

	go say(channel)

	log.Println("[MAIN] executando main...")
	time.Sleep(2 * time.Second)

	log.Println("[MAIN]", <- channel)
	log.Println("[MAIN]", <- channel)

	time.Sleep(1 * time.Second)
	log.Println("[MAIN]", <- channel)
	time.Sleep(1 * time.Second)
}

func say(c chan string) {
	c <- "buffer message 1"
	c <- "buffer message 2"
	log.Println("[GOROUTINE] executando thread...")
	c <- "buffer message 3"
	log.Println("[GOROUTINE] buffer liberado...")
}
