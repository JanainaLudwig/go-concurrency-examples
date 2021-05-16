package main

import (
	"sync"
)

type Monkey bool

type MonkeysList struct {
	List      []Monkey
	mutexList sync.Mutex
}

func (m *MonkeysList) GetList() []Monkey {
	m.mutexList.Lock()
	defer m.mutexList.Unlock()

	return m.List
}

func (m *MonkeysList) AddToList(monkey Monkey)  {
	m.mutexList.Lock()
	m.List = append(m.List, monkey)
	m.mutexList.Unlock()
}

func (m *MonkeysList) RemoveFromList(qtd int)  {
	m.mutexList.Lock()
	m.List = m.List[: len(m.List) - qtd]
	m.mutexList.Unlock()
}
