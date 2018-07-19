package main_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var timeJa int64 = 10

type r interface {
	myRand() int
}

type Rang struct {
	Min int
	Max int
}

type RangMock struct {
	Mock int
}

func (x Rang) myRand() int {
	return random(x.Min, x.Max)
}

func (x RangMock) myRand() int {
	return x.Mock
}

func random(min, max int) int {
	timeJa = timeJa + 1
	rand.Seed(time.Now().Unix() + timeJa)
	return rand.Intn(max-min) + min
}

func Sum(firstNum, secNum r) int {
	fmt.Println(firstNum)
	a := firstNum.myRand()
	b := secNum.myRand()

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return a + b
}

func TestSum(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		mockFirstNum := RangMock{
			Mock: 10,
		}
		mockSecNum := RangMock{
			Mock: 60,
		}

		got := Sum(mockFirstNum, mockSecNum)

		want := 70

		if got != want {
			t.Errorf("want %d, But got %d", want, got)
		}
	})
}
