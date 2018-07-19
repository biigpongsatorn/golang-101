package main

import (
	"fmt"
	"math/rand"
	"time"
)

var timeJa int64 = 10

func random() int {
	min := 0
	max := 100
	timeJa = timeJa + 1
	rand.Seed(time.Now().Unix() + timeJa)
	return rand.Intn(max-min) + min
}

func Sum() int {
	a := random()
	b := random()
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return a + b
}

func main() {
	fmt.Println(Sum())
}
