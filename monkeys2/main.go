package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	Limit int
	Added int
}

func (c *Counter) Increment(val int) {
	c.Lock()
	c.Added += val
	c.Unlock()
}

func (c *Counter) Finished() bool {
	c.Lock()
	defer c.Unlock()

	return c.Added == c.Limit
}


func main()  {
	//limit := 50
	//passed := 0

	leftChannel := make(chan Monkey, 5)
	rightChannel := make(chan Monkey, 5)

	monkeysOnLeft := MonkeysList{}
	monkeysOnRight := MonkeysList{}

	counter := Counter{
		Limit: 10,
	}
	bridge := Bridge{
		TimeToCross: time.Second,
	}

	bridge.timer = *time.NewTimer(time.Duration(0))

	log.Println("Total: ", counter.Limit)

	go generateMonkeys(leftChannel, 500)
	go generateMonkeys(rightChannel, 1000)

	for {
		select {
		case monkey := <-leftChannel: // Randomly add monkeys to the left side
			monkeysOnLeft.AddToList(monkey)
			log.Println("l")
			//log.Println("LEFT", len(monkeysOnLeft.GetList()))
		case monkey := <-rightChannel: // Randomly add monkeys to the right side
			log.Println("r")
			monkeysOnRight.AddToList(monkey)
			//log.Println("RIGHT", len(monkeysOnRight.GetList()))
		default:
			if counter.Finished() {
				log.Println("Finished after passing", counter.Limit, "monkeys")
				return
			}

			waitingRight := len(monkeysOnRight.GetList())
			waitingLeft := len(monkeysOnLeft.GetList())

			if waitingLeft > 0 {
				passed := bridge.Pass(LEFT)
				if passed {
					log.Println(waitingLeft, "LEFT -> _")
					monkeysOnLeft.RemoveFromList(waitingLeft)
					counter.Increment(waitingLeft)
				}
				continue
			}

			if waitingRight > 0 {
				passed := bridge.Pass(RIGHT)
				if passed {
					log.Println(waitingRight, "_ -> RIGHT")
					monkeysOnRight.RemoveFromList(waitingRight)
					counter.Increment(waitingRight)
				}
			}
		}
	}
}

func generateMonkeys(monkeys chan<- Monkey, delay int) {
	for  {
		//Random sleep
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(3000) + delay
		time.Sleep(time.Duration(n)*time.Millisecond)

		monkeys <- true
	}
}
