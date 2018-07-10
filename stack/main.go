package main

import (
	"fmt"
)

type stack []int

func main() {
	var value stack
	value.push(5)
	fmt.Println("result 1 : ", value)
	value.push(3)
	fmt.Println("result 2 : ", value)
	value.pop()
	fmt.Println("result 3 : ", value)
}

func (s *stack) push(a int) {
	*s = append(*s, a)
}

func (s *stack) pop() error {
	if len(*s) > 0 {
		a := *s
		_, a = a[len(a)-1], a[:len(a)-1]
		*s = a
		return nil
	} else {
		return nil
	}
}
