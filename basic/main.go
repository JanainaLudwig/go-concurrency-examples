package main

import (
	"errors"
	"log"
)

type Example struct {
	Value string
	Breda int
	E Example2
}

type Example2 struct {
	Value string
	Breda int
}

func (e *Example) getValue() (string, error) {
	defer log.Println("Defer é colocado em uma pilha de execução")

	if e.Value == "erro" {
		return "", errors.New("um erro aconteceu")
	}

	return e.Value, nil
}

func (e *Example) PrintValue() {
	val, err := e.getValue()

	if err != nil {
		log.Println("[ERRO]", err)
		return
	}

	log.Println("[VALOR]", val)
}

func main()  {
	example := Example{Value: "Valor"}

	example.PrintValue()

	example.Value = "erro"

	example.PrintValue()

	for i := 0; i < 5; i++ {
		log.Println("Valor de i:", i)
	}
}

func oi()  {
	log.Println("oi")
}