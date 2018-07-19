package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%#v\n", wordCount("I am Biig Pongsatorn EiEi , I am not you EiEi"))
}

func wordCount(s string) map[string]int {
	result := make(map[string]int)
	ar := strings.Fields(s)
	for _, a := range ar {
		result[a]++
	}
	return result
}
