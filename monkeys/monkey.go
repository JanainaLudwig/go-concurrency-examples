package main

import "sync"

type Monkey bool

type Monkeys struct {
	List []Monkey
	mutex sync.Mutex
}

func (m *Monkeys) Get() []Monkey {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.List
}

func (m *Monkeys) Add()  {
	m.mutex.Lock()
	m.List = append(m.List, true)
	m.mutex.Unlock()
}
