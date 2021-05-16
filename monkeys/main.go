package main

import (
	"math/rand"
	"time"
)


func main()  {
	//limit := 50
	//passed := 0

	monkeysLeft := Monkeys{}
	monkeysRight := Monkeys{}

	//var bridge sync.Mutex

	go addMonkey(&monkeysLeft)
	go addMonkey(&monkeysRight)


}

func addMonkey(monkeys *Monkeys)  {
	randomSleep()
	monkeys.Add()
}


func randomSleep()  {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)
	time.Sleep(time.Duration(n)*time.Second)
}