package main

import (
	"log"
	"time"
)

func main()  {
	channel := make(chan string)
	quit := make(chan string)

	go thread(channel, 200 * time.Microsecond)
	go thread(quit, 1000 * time.Microsecond)

	for {
		select {
		case val := <- channel:
			log.Println("[channel]", val)
		case val := <- quit:
			log.Println("[channel]", val, "QUITING...")
			return
		default:
			log.Println("waiting")
		}
	}
}

func thread(c chan string, duration time.Duration) {
	time.Sleep(duration)
	c <- "message from thread"
}