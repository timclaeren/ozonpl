package main

import (
	"fmt"
	"math"
)

var k float64 = 129
var p float64 = -6
var v float64 = -6

func T() float64 {
	n := W() / 6
	return n
}

func W() float64 {
	n := math.Sqrt(k / M())
	return n
}

func M() float64 {
	n := p * v
	return n
}

func main() {
	fmt.Println(T())
}
