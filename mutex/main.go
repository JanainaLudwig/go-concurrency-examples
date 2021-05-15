package main

import (
	"log"
	"sync"
)

func main()  {
	var wg sync.WaitGroup

	data := Data{value: 10}

	wg.Add(3)

	go add(&wg, &data, 5)
	go multiply(&wg, &data, 10)
	go add(&wg, &data, 10)

	wg.Wait()

	log.Println("Final result:", data.GetValue())
}

func add(wg *sync.WaitGroup, d *Data, value int) {
	defer wg.Done()

	d.SetValue(d.GetValue() + value)
}

func multiply(wg *sync.WaitGroup, d *Data, value int) {
	defer wg.Done()

	d.SetValue(d.GetValue() * value)
}