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
	bridgeMutex sync.Mutex
	TimeToCross time.Duration
	CurrentDirection Direction
	Debug bool
}

func (c *Bridge) Pass(d Direction) bool {

	c.bridgeMutex.Lock()

	//log.Println("Current", getDirection(c.CurrentDirection), "  Desired", getDirection(d))
	if c.CurrentDirection != EMPTY && d != c.CurrentDirection {
		c.bridgeMutex.Unlock()
		return false
	}

	c.CurrentDirection = d

	go func() {
		if c.Debug {
			log.Println("crossing... ", getDirection(d))
		}

		time.Sleep(c.TimeToCross)

		if c.Debug {
			log.Println("crossed...")
		}

		c.CurrentDirection = EMPTY
		c.bridgeMutex.Unlock()
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