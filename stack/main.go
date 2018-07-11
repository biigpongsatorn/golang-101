package main

import (
	"errors"
	"fmt"
)

type stack []int

func main() {
	var value stack
	value.push(5)
	fmt.Println("push(5) : ", value)
	value.push(3)
	fmt.Println("push(3) : ", value)
	value.pop()
	fmt.Println("pop() : ", value)
	value.push(7)
	fmt.Println("push(7) : ", value)
	value.push(8)
	fmt.Println("push(8) : ", value)
	value.push(1)
	fmt.Println("push(1) : ", value)
	value.pop()
	fmt.Println("pop() : ", value)
	value.pop()
	fmt.Println("pop() : ", value)
	value.pop()
	fmt.Println("pop() : ", value)
	value.pop()
	fmt.Println("pop() : ", value)
	err := value.pop()
	if err != nil {
		fmt.Println("pop() : ", err)
	} else {
		fmt.Println("pop() : ", value)
	}
}

func (s *stack) push(a int) {
	*s = append(*s, a)
}

func (s *stack) pop() error {
	if len(*s) > 0 {
		a := *s
		a = a[:len(a)-1]
		*s = a
		return nil
	}
	return errors.New("error na ja")
}
