package main

import (
	"log"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup

	wg.Add(1)
	go sayWorld(&wg)

	log.Println("Hello")

	wg.Wait()
}

func sayWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	log.Println("World")
}