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
	directionMutex sync.Mutex
	timer time.Timer
}
//
//func (c *Bridge) CanPass() bool {
//	c.directionMutex.Lock()
//	defer c.directionMutex.Unlock()
//
//	return c.CurrentDirection == EMPTY
//}
//
//func (c *Bridge) Pass(d Direction) bool {
//
//	c.directionMutex.Lock()
//	defer c.directionMutex.Unlock()
//
//	if c.CurrentDirection != EMPTY && d != c.CurrentDirection {
//		return false
//	}
//
//	defer func() {
//		c.bridgeMutex.Lock()
//		c.timer.Reset(c.TimeToCross)
//		c.CurrentDirection = d
//
//		<- c.timer.C
//		c.bridgeMutex.Unlock()
//		c.CurrentDirection = EMPTY
//	}()
//
//	return true
//}

func (c *Bridge) Pass(d Direction) bool {

	c.bridgeMutex.Lock()
	//defer c.bridgeMutex.Unlock()

	if c.CurrentDirection != EMPTY && d != c.CurrentDirection {
		log.Println(d, c.CurrentDirection)
		c.bridgeMutex.Unlock()
		return false
	}

	go func() {
		time.Sleep(c.TimeToCross)
		c.bridgeMutex.Unlock()
	}()

	return true
}
