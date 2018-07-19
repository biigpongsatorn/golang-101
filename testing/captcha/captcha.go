package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	captchaString, result := generateCaptcha()
	fmt.Printf("%s = %d", captchaString, result)
}

var timeJa int64 = 10

func random(min, max int) int {
	timeJa = timeJa + 1
	rand.Seed(time.Now().Unix() + timeJa)
	return rand.Intn(max-min) + min
}

func generateCaptcha() (string, int) {
	firstCaptcha := random(10, 99)
	secondCaptcha := random(10, 99)

	result := 0
	operator := ""
	switch operationRandom := random(0, 3); operationRandom {
	case 0:
		result = firstCaptcha + secondCaptcha
		operator = "+"
	case 1:
		result = firstCaptcha - secondCaptcha
		operator = "-"
	case 2:
		result = firstCaptcha * secondCaptcha
		operator = "*"
	default:
		result = firstCaptcha / secondCaptcha
		operator = "/"
	}

	captchaString := fmt.Sprintf("%d %s %d", firstCaptcha, operator, secondCaptcha)
	return captchaString, result
}
