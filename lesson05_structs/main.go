package main

import (
	"fmt"
)

type IntStack struct {
	stack []int
}

func (i *IntStack) Push(num int) {
	i.stack = append(i.stack, num)
}

func (i *IntStack) Pop() int {
	index := len(i.stack) - 1
	if index < 0 {
		return 0
	}

	res := i.stack[index]
	i.stack = i.stack[:index]
	return res
}

func main() {
	s := IntStack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("expected 30, got %d\n", s.Pop())
	fmt.Printf("expected 20, got %d\n", s.Pop())
	fmt.Printf("expected 10, got %d\n", s.Pop())
	fmt.Printf("expected ?, got %d\n", s.Pop())
}
