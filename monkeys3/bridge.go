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
	Debug            bool
	debugId          int
	esquerda 		 sync.WaitGroup
	direita 		 sync.WaitGroup
}

func (c *Bridge) PassRight(wg *sync.WaitGroup, d Direction, qtd int) {
	c.esquerda.Wait()

	c.mutex.Lock()
	c.debugId++
	i := c.debugId
	c.direita.Add(1)
	c.mutex.Unlock()

	log.Printf("[ %02d ] START %v (%v)", i, getDirection(d), qtd)

	wg.Add(1)
	go func() {
		defer c.direita.Done()
		defer wg.Done()

		time.Sleep(c.TimeToCross)

		log.Printf("[ %02d ] END", i)
	}()
}

func (c *Bridge) PassLeft(wg *sync.WaitGroup, d Direction, qtd int) {
	c.direita.Wait()

	c.mutex.Lock()
	c.debugId++
	i := c.debugId
	c.esquerda.Add(1)
	c.mutex.Unlock()

	log.Printf("[ %02d ] START %v (%v)", i, getDirection(d), qtd)

	wg.Add(1)
	go func() {
		defer c.esquerda.Done()
		defer wg.Done()

		time.Sleep(c.TimeToCross)

		log.Printf("[ %02d ] END", i)
	}()

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