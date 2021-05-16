package main

import "sync"

type Monkey bool

type MonkeysCount struct {
	List      []Monkey
	Limit     int
	mutexList sync.Mutex
}

func (m *MonkeysCount) GetList() []Monkey {
	m.mutexList.Lock()
	defer m.mutexList.Unlock()

	return m.List
}

func (m *MonkeysCount) AddToList(monkey Monkey)  {
	m.mutexList.Lock()
	m.List = append(m.List, monkey)
	m.mutexList.Unlock()
}

func (m *MonkeysCount) LimitReached() bool {
	return m.Limit == len(m.GetList())
}
