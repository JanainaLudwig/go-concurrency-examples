package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	MinMonkeysGenerated int
	Added               int
}

func (c *Counter) Increment(val int) {
	c.Lock()
	c.Added += val
	c.Unlock()
}

func (c *Counter) Finished() bool {
	c.Lock()
	defer c.Unlock()

	return c.Added >= c.MinMonkeysGenerated
}

func main()  {
	leftChannel := make(chan Monkey, 10)
	rightChannel := make(chan Monkey, 10)

	monkeysOnLeft := MonkeysList{}
	monkeysOnRight := MonkeysList{}

	counter := Counter{
		MinMonkeysGenerated: 20,
	}
	bridge := Bridge{
		TimeToCross: 5 * time.Second,
		Debug: false,
	}

	log.Println("Total: ", counter.MinMonkeysGenerated)

	go generateMonkeys(leftChannel, 500)
	go generateMonkeys(rightChannel, 300)

	for {
		select {
		case monkey := <-leftChannel: // Randomly add monkeys to the left side
			monkeysOnLeft.AddToList(monkey)
			//log.Println("LEFT", len(monkeysOnLeft.GetList()))
		case monkey := <-rightChannel: // Randomly add monkeys to the right side
			monkeysOnRight.AddToList(monkey)
			//log.Println("RIGHT", len(monkeysOnRight.GetList()))
		default:
			if counter.Finished() {
				log.Println("Finished after passing", counter.MinMonkeysGenerated, "monkeys")
				return
			}

			waitingRight := len(monkeysOnRight.GetList())
			waitingLeft := len(monkeysOnLeft.GetList())


			if waitingRight > 0 {
				passed := bridge.Pass(RIGHT)
				if passed {
					log.Println(waitingRight, "_ <- RIGHT")
					monkeysOnRight.RemoveFromList(waitingRight)
					counter.Increment(waitingRight)
				}
			}

			if waitingLeft > 0 {
				passed := bridge.Pass(LEFT)
				if passed {
					log.Println(waitingLeft, "LEFT -> _")
					monkeysOnLeft.RemoveFromList(waitingLeft)
					counter.Increment(waitingLeft)
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
