// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// var timeJa int64 = 10

// type r interface {
// 	rand() int
// }

// type Rang struct {
// 	Min int
// 	Max int
// }

// func (x Rang) rand() int {
// 	return random(x.Min, x.Max)
// }

// func random(min, max int) int {
// 	timeJa = timeJa + 1
// 	rand.Seed(time.Now().Unix() + timeJa)
// 	return rand.Intn(max-min) + min
// }

// func Sum(firstNum, secNum r) int {
// 	a := firstNum.rand()
// 	b := secNum.rand()

// 	fmt.Println("a = ", a)
// 	fmt.Println("b = ", b)

// 	return a + b
// }
// func main() {
// 	got := Sum(Rang{
// 		Min: 0,
// 		Max: 50,
// 	}, Rang{
// 		Min: 50,
// 		Max: 100,
// 	})
// 	fmt.Println(got)
// }
