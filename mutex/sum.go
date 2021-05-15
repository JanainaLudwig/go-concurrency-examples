package main

import (
	"sync"
)

type Data struct {
	value int
	mutex sync.Mutex
}

func (d *Data) SetValue(value int) {
	d.mutex.Lock()

	d.value = value

	d.mutex.Unlock()
}

func (d *Data) GetValue() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.value
}