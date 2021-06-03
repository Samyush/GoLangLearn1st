package main

import (
	"fmt"
)

var s string

func main() {
	fmt.Println("first code samyush 😍")

	var x float64
	var y float64

	x = 1
	y = 2

	fmt.Printf("x= %v, type of %T\n", x, x)
	fmt.Printf("x= %v, type of %T\n", y, y)

	var mean float64
	mean = (x + y) / 2.0

	meanNew()

	conditional()

	switchCondition()

	forLoop()

	goString()

	mapData()

	fmt.Printf("result: %v, type of %T\n", mean, mean)
}

// implementation of numbers and assignments
func meanNew() {
	a := 1.0
	b := 2.0

	mean := (a + b) / 2
	fmt.Printf("new result: %v, type of %T\n", mean, mean)
}

// implementation of if conditional
func conditional() {
	con := 1

	if con > 5 {
		fmt.Println("is greater")
	} else {
		fmt.Println("is small")
	}

	// the below code here gives the value of frac and applied itself in the conditional
	if frac := 1.0 / 3.0; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}

// implementation of switch conditional
func switchCondition() {
	cons := 2

	switch cons {
	case 1:
		fmt.Println("is one")
	case 2:
		fmt.Println("is two")
	case 3:
		fmt.Println("is three")
	default:
		fmt.Println("is many")
	}
}

// implementing of for loop
func forLoop() {
	for data1 := 0; data1 < 10; data1++ {
		if data1 == -1 {
			break
		}
		println(data1)
	}
}

// working with strings
func goString() {
	book := "The colour of magic"
	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// unit8 is a byte

	// using slice (start, end)
	fmt.Println(book[4:11])

	// Double value range, igonors index by using _
	for _, name := range book {
		fmt.Println(name)
	}
}

func mapData() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}

	// NUMBER OF ITEMS
	fmt.Println(len(stocks))

	// GET A VALUE
	fmt.Println(stocks["MSFT"])

	// GET ZERO VALUE IF NOT FOUND
	fmt.Println(stocks["TSLA"])

	// use two value to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// set
	stocks["TSLA"] = 322.12

	// Double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)

	}
}
