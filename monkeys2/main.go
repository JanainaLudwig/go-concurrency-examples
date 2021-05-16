package main

import (
	"log"
	"math/rand"
	"time"
)


func main()  {
	//limit := 50
	//passed := 0

	monkeysLeft := make(chan Monkey, 10)
	monkeysRight := make(chan Monkey, 10)

	//var bridge sync.Mutex
	counter := MonkeysCount{
		Limit: 10,
	}

	log.Println("Total: ", counter.Limit)

	go addMonkey(monkeysLeft)
	go addMonkey(monkeysRight)

	for {
		select {
		case monkey := <- monkeysLeft:
			log.Println("LEFT -> _")
			time.Sleep(20 * time.Millisecond)
			counter.AddToList(monkey)
		case monkey := <- monkeysRight:
			log.Println("_ <- RIGHT")
			time.Sleep(20 * time.Millisecond)
			counter.AddToList(monkey)
		default:
			if counter.LimitReached() {
				log.Println("Finished after passing", len(counter.GetList()), "monkeys")
				return
			}
		}
	}

}

func addMonkey(monkeys chan<- Monkey) {
	for  {
		//Random sleep
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000) + 200
		time.Sleep(time.Duration(n)*time.Millisecond)

		monkeys <- true
	}
}


//func randomSleep() int {
//	rand.Seed(time.Now().UnixNano())
//	n := rand.Intn(1000) + 200
//	time.Sleep(time.Duration(n)*time.Millisecond)
//
//	return n
//}