package main

import (
	"log"

	"github.com/meesooqa/docker-bin-go/calc"
)

func main() {
	log.Println("Hello World")

	a := 120
	b := 10
	c, _ := calc.Divide(float64(a), float64(b))

	// 120 divided by 10 equals 12
	log.Printf("%d divided by %d equals %f\n", a, b, c)
}
