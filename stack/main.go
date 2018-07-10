package main

type stack []int

func main() {
	var value stack
	value.push(5)
	value.pop()
}

func (s stack) push(a int) stack {

	return s
}

func (s stack) pop() (int, error) {
	return 0, nil
}
