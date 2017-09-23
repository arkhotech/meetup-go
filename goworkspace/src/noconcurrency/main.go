package main

import (
	"fmt"
	"math"
)

func main() {

	abc := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", ""}
	lenAbc := len(abc)
	
	for i := 0; i < lenAbc; i++ {
		imprimirRuna(abc[i])
	}

}

func imprimirRuna(letra string) {

	if letra != "" {
		fmt.Printf("%s", letra)
	}

	//Demorar
	result := 0.0
	for i := 0; i < 10000000; i++ {
		result += math.Pi * math.Sin(float64(i))
	}

}
