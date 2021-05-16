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
	wg := sync.WaitGroup{}
	leftChannel := make(chan Monkey, 10)
	rightChannel := make(chan Monkey, 10)

	monkeysOnLeft := MonkeysList{}
	monkeysOnRight := MonkeysList{}

	counter := Counter{
		MinMonkeysGenerated: 30,
	}
	bridge := Bridge{
		TimeToCross: 1 * time.Second,
		Debug: true,
	}

	log.Println("Total: ", counter.MinMonkeysGenerated)

	go generateMonkeys(leftChannel, 1000)
	go generateMonkeys(rightChannel, 1000)

	for {
		select {
		case monkey := <-leftChannel: // Randomly add monkeys to the left side
			monkeysOnLeft.AddToList(monkey)
		case monkey := <-rightChannel: // Randomly add monkeys to the right side
			monkeysOnRight.AddToList(monkey)
		default:
			if counter.Finished() {
				wg.Wait()
				log.Println("Finished after passing", counter.MinMonkeysGenerated, "monkeys")
				return
			}

			waitingRight := len(monkeysOnRight.GetList())
			waitingLeft := len(monkeysOnLeft.GetList())

			if waitingRight > 0 {
				passed := bridge.Pass(&wg, RIGHT)
				if passed {
					monkeysOnRight.RemoveFromList(waitingRight)
					counter.Increment(waitingRight)
				}
			}

			if waitingLeft > 0 {
				passed := bridge.Pass(&wg, LEFT)
				if passed {
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
