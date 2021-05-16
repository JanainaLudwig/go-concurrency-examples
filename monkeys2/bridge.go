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

func (c *Bridge) Pass(d Direction) bool {
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

	go func() {
		if c.Debug {
			log.Println("[", i , "]", "start", getDirection(d))
		}

		time.Sleep(c.TimeToCross)

		if c.Debug {
			log.Println("[", i , "]", "end", getDirection(d))
		}

		c.mutex.Lock()
		c.PassingCount--

		if c.PassingCount == 0 {
			c.CurrentDirection = EMPTY
		}

		c.mutex.Unlock()
	}()

	return true
}

func getDirection(d Direction) string {
	switch d {
	case RIGHT:
		return "Right"
	case LEFT:
		return "Left"
	default:
		return "Empty"
	}
}