package main

import (
	"log"
	"time"
)

func main()  {
	go sayWorld()

	log.Println("Hello")

	time.Sleep(3 * time.Second)
}

func sayWorld() {
	time.Sleep(time.Second)
	log.Println("World")
}