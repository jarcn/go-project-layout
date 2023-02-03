package test

import (
	"fmt"
	"testing"
)

type Person1 struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person1 {
	return &Person1{
		name: name,
		age:  age,
	}
}

func (p Person1) Name() string {
	return p.name
}

func TestPerson1(t *testing.T) {
	p := NewPerson("chenjia", 31)
	fmt.Println(p.Name())
}
