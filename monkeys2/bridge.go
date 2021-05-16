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
	directionMutex sync.Mutex
	TimeToCross time.Duration
	CurrentDirection Direction
	Debug bool
}

func (c *Bridge) Pass(d Direction) bool {
	c.directionMutex.Lock()

	//log.Println("Current", getDirection(c.CurrentDirection), "  Desired", getDirection(d))
	if c.CurrentDirection != EMPTY && d != c.CurrentDirection {
		c.directionMutex.Unlock()
		return false
	}

	c.CurrentDirection = d
	c.directionMutex.Unlock()

	c.bridgeMutex.Lock()
	go func() {
		if c.Debug {
			log.Println("crossing... ", getDirection(d))
		}

		time.Sleep(c.TimeToCross)

		if c.Debug {
			log.Println("crossed...", getDirection(d))
		}

		c.directionMutex.Lock()
		c.CurrentDirection = EMPTY
		c.directionMutex.Unlock()

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