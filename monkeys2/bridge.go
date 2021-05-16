package main

import (
	"log"
	"sync"
	"time"
)

type Direction int8

const (
	EMPTY = iota
	LEFT
	RIGHT
)

type Bridge struct {
	mutex            sync.Mutex
	TimeToCross      time.Duration
	CurrentDirection Direction
	PassingCount     int
	Debug            bool
	debugId          int
}

func (c *Bridge) Pass(wg *sync.WaitGroup, d Direction) bool {
	c.mutex.Lock()

	if c.CurrentDirection != EMPTY && d != c.CurrentDirection {
		c.mutex.Unlock()
		return false
	}

	c.CurrentDirection = d
	c.PassingCount++

	c.mutex.Unlock()

	c.debugId++
	i := c.debugId

	wg.Add(1)
	go func() {
		if c.Debug {
			log.Printf("[ %02d ] START %v", i, getDirection(d))
		}

		time.Sleep(c.TimeToCross)

		if c.Debug {
			log.Printf("[ %02d ] END", i)
		}

		c.mutex.Lock()
		c.PassingCount--

		if c.PassingCount == 0 {
			c.CurrentDirection = EMPTY
		}

		c.mutex.Unlock()
		wg.Done()
	}()

	return true
}

func getDirection(d Direction) string {
	switch d {
	case RIGHT:
		return "____ <- Right"
	case LEFT:
		return "Left -> _____"
	default:
		return "Empty"
	}
}