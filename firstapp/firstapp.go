package main

import (
	"fmt"

	"github.com/yabuking84/golang-experiments/firstapp/greeting"
)

func main() {
	number := 17

	// yo := number + 2

	number = 34
	div := 3

	// var asd rune = '\U+1F600'
	asd := '\U0001f603'
	a1 := '\U0001f0cf'

	greeting.GreetingText = "asdsadsa"

	var newNumner float64 = float64(number / div)
	// var newNumner float64 = float64(number) / float64(div)
	fmt.Println(newNumner)
	fmt.Println(string(asd))
	fmt.Println(string(a1))
	fmt.Printf("hey %T %v %v", div, greeting.GreetingText, greeting.PI)
	fmt.Println()
}
